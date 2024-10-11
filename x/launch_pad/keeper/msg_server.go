package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/types"
)

type msgServer struct {
	keeper *Keeper
}

func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

func (server msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	poolAddress := GenerateNewAddress("", "pool", sdk.Uint64ToBigEndian(1))

	_, err := server.keeper.CreateNewPool(goCtx, msg, poolAddress)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreatePoolResponse{}, nil
}

func (server msgServer) SwapExactIn(goCtx context.Context, msg *types.MsgSwapExactIn) (*types.MsgSwapExactInResponse, error) {
	return &types.MsgSwapExactInResponse{}, nil
}

func (server msgServer) SwapExactOut(goCtx context.Context, msg *types.MsgSwapExactOut) (*types.MsgSwapExactOutResponse, error) {
	return &types.MsgSwapExactOutResponse{}, nil
}
