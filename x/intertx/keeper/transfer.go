package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	"time"
)

func (k Keeper) TransferAssets(
	goCtx context.Context,
	id uint64,
	sender sdk.AccAddress, receiver string,
	token sdk.Coin, memo string) error {
	k.Logger.Info(fmt.Sprintf("TransferAssets, id=%v, sender=%v, token=%v, memo=%v", id, sender.String(), token.String(), memo))
	ctx := sdk.UnwrapSDKContext(goCtx)
	port := k.transferKeeper.GetPort(ctx)
	channel := k.channelKeeper.GetAllChannelsWithPortPrefix(ctx, port)

	_, err := k.transferKeeper.Transfer(goCtx, &transfertypes.MsgTransfer{
		SourcePort:    port,
		SourceChannel: channel[0].ChannelId,
		Sender:        sender.String(),
		Receiver:      receiver,
		Token:         token,
		TimeoutHeight: ibcclienttypes.Height{
			RevisionHeight: 0,
			RevisionNumber: 0,
		},
		TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano() + 5*time.Minute.Nanoseconds()),
		Memo:             memo,
	})
	if err != nil {
		return err
	}

	return nil
}
