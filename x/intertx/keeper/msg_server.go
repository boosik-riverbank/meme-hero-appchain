package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	"github.com/cosmos/interchain-security/v6/x/intertx/types"
)

type msgServer struct {
	keeper *Keeper
}

func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

func (m msgServer) CreateButler(goCtx context.Context, msg *types.MsgCreateButler) (*types.MsgCreateButlerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: check controller address
	b, err := m.keeper.IsExist(goCtx, msg.TargetChainId)
	if err != nil {
		return nil, err
	}

	if b {
		return nil, fmt.Errorf("already exist butler id")
	}

	if msg.FromAddress == "" {
		return nil, fmt.Errorf("from address is not set")
	}

	_, err = sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, fmt.Errorf("cannot parse from address")
	}

	version := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: msg.ConnectionId,
		HostConnectionId:       msg.ConnectionId,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))

	moduleAddress := m.keeper.GetModuleAcc(ctx).String()
	err = m.keeper.IcaControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, moduleAddress, version)
	if err != nil {
		return nil, err
	}

	portId, err := icatypes.NewControllerPortID(moduleAddress)
	if err != nil {
		return nil, err
	}

	_, err = m.keeper.CreateButler(goCtx, msg, portId)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateButlerResponse{}, nil
}

func (m msgServer) SubmitTx(ctx context.Context, tx *types.MsgSubmitTx) (*types.MsgSubmitTxResponse, error) {
	return nil, nil
}

func (m msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	b, err := m.keeper.GetButler(goCtx, msg.TargetChainId)
	if err != nil {
		return nil, err
	}

	senderAcc, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	nextId, err := m.keeper.GetNextQueueId(ctx)
	if err != nil {
		return nil, err
	}

	if err = m.keeper.AddQueue(ctx, nextId, b.TargetChainId, *msg.Token, *msg.Pair); err != nil {
		return nil, err
	}

	memo := fmt.Sprintf("%v:%s:%v", nextId, types.MEMO_TRANSFER_TOKEN, msg.TargetChainId)
	err = m.keeper.TransferAssets(goCtx, msg.Id, senderAcc, b.HostAddress, *msg.Token, memo)
	if err != nil {
		return nil, err
	}

	err = m.keeper.SetNextQueueId(ctx, nextId+1)
	if err != nil {
		return nil, err
	}

	return &types.MsgTransferResponse{}, nil
}
