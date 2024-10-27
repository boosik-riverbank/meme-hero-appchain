package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	cosmossdk_io_math "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	transferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcchannelkeeper "github.com/cosmos/ibc-go/v8/modules/core/04-channel/keeper"
	"github.com/cosmos/interchain-security/v6/x/intertx/types"
	"github.com/osmosis-labs/osmosis/osmomath"
	osmosistypes "github.com/osmosis-labs/osmosis/v26/x/gamm/pool-models/balancer"
	"time"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.Codec

	paramSpace paramtypes.Subspace

	authKeeper          types.AccountKeeper
	scopedKeeper        capabilitykeeper.ScopedKeeper
	IcaControllerKeeper icacontrollerkeeper.Keeper
	transferKeeper      transferkeeper.Keeper
	hooks               types.ICAHooks
	channelKeeper       ibcchannelkeeper.Keeper

	butlers     collections.Map[uint64, types.Butler]
	queue       collections.Map[uint64, types.Queue]
	nextQueueId collections.Item[uint64]

	Logger       log.Logger
	eventService runtime.EventService
}

func NewKeeper(
	storeService store.KVStoreService,
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	ak types.AccountKeeper,
	iaKeeper icacontrollerkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	paramStore paramtypes.Subspace,
	channelKeeper ibcchannelkeeper.Keeper,
	transferKeeper transferkeeper.Keeper,
	eventService runtime.EventService,
	logger log.Logger) Keeper {
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	sb := collections.NewSchemaBuilder(storeService)
	return Keeper{
		cdc:                 cdc,
		storeKey:            storeKey,
		authKeeper:          ak,
		scopedKeeper:        scopedKeeper,
		paramSpace:          paramStore,
		IcaControllerKeeper: iaKeeper,
		channelKeeper:       channelKeeper,
		transferKeeper:      transferKeeper,
		butlers:             collections.NewMap(sb, types.KeyPrefixButlers, "butlers", collections.Uint64Key, codec.CollValue[types.Butler](cdc)),
		queue:               collections.NewMap(sb, types.KeyPrefixQueue, "queue", collections.Uint64Key, codec.CollValue[types.Queue](cdc)),
		nextQueueId:         collections.NewItem(sb, types.KeyPrefixNextQueueId, "next_queue_id", collections.Uint64Value),
		eventService:        eventService,
		Logger:              logger,
	}
}

func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

func (k *Keeper) SetHooks(hook types.ICAHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set ICA hooks twice")
	}

	k.hooks = hook

	return k
}

func (k Keeper) GetModuleAcc(ctx sdk.Context) sdk.AccAddress {
	return k.authKeeper.GetModuleAddress(types.ModuleName)
}

func (k Keeper) SendCreateNewPoolTx(ctx sdk.Context, targetChainId uint64, token, pair sdk.Coin) error {
	b, err := k.GetButler(ctx, targetChainId)
	if err != nil {
		return err
	}

	moduleAddress := k.GetModuleAcc(ctx)
	transferPort := k.transferKeeper.GetPort(ctx)
	//transferChannel := k.channelKeeper.Connection
	tokenIbcDenom := k.GetIBCHashDenom(transferPort, "channel-0", token.Denom)
	pairIbcDenom := k.GetIBCHashDenom(transferPort, "channel-0", pair.Denom)
	var msgs []sdk.Msg
	k.Logger.Info(fmt.Sprintf("MsgCreateBalancerPool: portId=%v, chanId=%v, tokenIbcDenom=%v, pairIbcDenom=%v", b.IcaConnectionInfo.PortId, b.IcaConnectionInfo.ChannelId, tokenIbcDenom, pairIbcDenom))
	msgs = append(msgs, &osmosistypes.MsgCreateBalancerPool{
		Sender: b.HostAddress,
		PoolParams: &osmosistypes.PoolParams{
			SwapFee: osmomath.MustNewDecFromStr("0.025"),
			ExitFee: osmomath.MustNewDecFromStr("0"),
		},
		PoolAssets: []osmosistypes.PoolAsset{
			{
				Token:  sdk.NewCoin(tokenIbcDenom, token.Amount),
				Weight: cosmossdk_io_math.NewInt(1),
			},
			{
				Token:  sdk.NewCoin(pairIbcDenom, pair.Amount),
				Weight: cosmossdk_io_math.NewInt(1),
			},
		},
	})
	err = k.SendTx(ctx, moduleAddress.String(), b.IcaConnectionInfo.ConnectionId, msgs, 30_000)
	if err != nil {
		return err
	}

	return nil
}

// GetIBCHashDenom uses baseDenom and portId and channelId to create the appropriate IBCdenom.
func (k Keeper) GetIBCHashDenom(portId, chanId, baseDenom string) string {
	var path string

	if portId == "" || chanId == "" {
		path = ""
	} else {
		path = portId + "/" + chanId
	}

	denomTrace := transfertypes.DenomTrace{
		Path:      path,
		BaseDenom: baseDenom,
	}

	denomHash := denomTrace.IBCDenom()

	return denomHash
}

func (k Keeper) SendTx(
	ctx sdk.Context,
	owner,
	connectionId string,
	msgs []sdk.Msg,
	timeout uint64) error {
	data, err := icatypes.SerializeCosmosTx(k.cdc, msgs, icatypes.EncodingProtobuf)
	if err != nil {
		return err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to satisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + timeout*uint64(time.Minute.Nanoseconds())

	icaMsgServer := icacontrollerkeeper.NewMsgServerImpl(&k.IcaControllerKeeper)
	_, err = icaMsgServer.SendTx(ctx, &icacontrollertypes.MsgSendTx{
		Owner:           owner,
		ConnectionId:    connectionId,
		PacketData:      packetData,
		RelativeTimeout: timeoutTimestamp,
	})
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) AddQueue(ctx sdk.Context, id uint64, targetChainId uint64, token, pair sdk.Coin) error {
	has, err := k.queue.Has(ctx, id)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("queue %s already exists", id)
	}

	err = k.queue.Set(ctx, id, types.Queue{
		Token:         &token,
		Pair:          &pair,
		TargetChainId: targetChainId,
		Finished:      false,
	})
	return err
}

func (k Keeper) GetQueue(ctx sdk.Context, id uint64) (types.Queue, error) {
	return k.queue.Get(ctx, id)
}

func (k Keeper) GetNextQueueId(ctx sdk.Context) (uint64, error) {
	exist, err := k.nextQueueId.Has(ctx)
	if err != nil {
		return 0, err
	}
	if !exist {
		return 0, nil
	}

	return k.nextQueueId.Get(ctx)
}

func (k Keeper) SetNextQueueId(ctx sdk.Context, id uint64) error {
	return k.nextQueueId.Set(ctx, id)
}
