package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/interchain-security/v6/x/launchpad/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec

	paramSpace paramtypes.Subspace

	feeCollector sdk.AccAddress

	// Data
	count collections.Item[uint64]
	pools collections.Map[uint64, types.BondingCurvePool]

	// Keepers
	authKeeper types.AuthKeeper
	bankKeeper types.BankKeeper

	logger       log.Logger
	eventService runtime.EventService
}

func NewKeeper(
	storeService store.KVStoreService,
	cdc codec.BinaryCodec,
	authKeeper types.AuthKeeper,
	bankKeeper types.BankKeeper,
	eventService runtime.EventService,
	logger log.Logger) Keeper {
	sb := collections.NewSchemaBuilder(storeService)
	return Keeper{
		// feeCollector: feeCollector,
		authKeeper:   authKeeper,
		bankKeeper:   bankKeeper,
		count:        collections.NewItem(sb, types.KeyPrefixPoolCount, "count", collections.Uint64Value),
		pools:        collections.NewMap(sb, types.KeyPrefixPools, "pools", collections.Uint64Key, codec.CollValue[types.BondingCurvePool](cdc)),
		logger:       logger,
		eventService: eventService,
	}
}
