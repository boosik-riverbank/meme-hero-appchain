package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcaccounttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/cosmos/interchain-security/v6/x/intertx/types"
)

func (k *Keeper) HandleAckMsgData(ctx sdk.Context, packet channeltypes.Packet, msgData *sdk.MsgData) (string, error) {
	return "", nil
}

func (k *Keeper) HandleAckFail(ctx sdk.Context, packet channeltypes.Packet) error {
	var data ibcaccounttypes.InterchainAccountPacketData
	if err := ibcaccounttypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return fmt.Errorf("cannot unmarshal packet data: %s", err.Error())
	}

	if data.Data == nil {
		return fmt.Errorf("invalid data")
	}

	packetData, err := ibcaccounttypes.DeserializeCosmosTx(k.cdc, data.Data, icatypes.EncodingProtobuf)
	if err != nil {
		return err
	}

	switch packetData[0].(type) {
	// TODO : handle tx when osmosis pool creating failed
	}

	return nil
}

func (k Keeper) GetConnectionId(ctx sdk.Context, portId string) (string, error) {
	icas := k.IcaControllerKeeper.GetAllInterchainAccounts(ctx)
	for _, ica := range icas {
		if ica.PortId == portId {
			return ica.ConnectionId, nil
		}
	}
	return "", fmt.Errorf("portId %s has no associated connectionId", portId)
}

func (k Keeper) GetInterchainAccountAddress(ctx sdk.Context, connectionID, portID string) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	key := icatypes.KeyOwnerAccount(portID, connectionID)

	if !store.Has(key) {
		return "", false
	}

	return string(store.Get(key)), true
}

func (k Keeper) GetButlerPortId(ctx context.Context, portId string) (*types.Butler, error) {
	var butler *types.Butler
	err := k.butlers.Walk(ctx, nil, func(_ uint64, b types.Butler) (stop bool, err error) {
		if b.IcaConnectionInfo.PortId == portId {
			butler = &b
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return butler, nil
}
