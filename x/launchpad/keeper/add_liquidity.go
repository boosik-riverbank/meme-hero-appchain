package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	intertxkeeper "github.com/cosmos/interchain-security/v6/x/intertx/keeper"
	intertxtypes "github.com/cosmos/interchain-security/v6/x/intertx/types"
)

func (k Keeper) AddLiquidity(goCtx context.Context, sender sdk.AccAddress, targetChainId uint64, token sdk.Coin, pair sdk.Coin) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	butler, err := k.icaKeeper.GetButler(goCtx, targetChainId)
	if err != nil {
		return err
	}
	if butler == nil {
		return fmt.Errorf("butler not found")
	}

	if butler.IcaConnectionInfo.PortId == "" {
		return fmt.Errorf("port id not set")
	}

	interTxMsgServer := intertxkeeper.NewMsgServerImpl(&k.icaKeeper)
	_, err = interTxMsgServer.Transfer(ctx, &intertxtypes.MsgTransfer{
		Id:            1,
		TargetChainId: targetChainId,
		Sender:        sender.String(),
		Token:         &token,
		Pair:          &pair,
		Memo:          "",
	})
	return err
}
