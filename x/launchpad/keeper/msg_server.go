package keeper

import (
	"context"
	"cosmossdk.io/core/event"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launchpad/types"
	"strconv"
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
	pool, err := server.keeper.CreateNewPool(goCtx, msg)
	if err != nil {
		return nil, err
	}

	server.keeper.logger.Info(fmt.Sprintf("Create new pool: %v", pool.String()))
	server.keeper.logger.Info(fmt.Sprintf("Event Service: %v", server.keeper.eventService))

	tokenAmount, pairAmount, err := server.keeper.GetPoolAmount(goCtx, pool.Id)
	if err != nil {
		return nil, err
	}

	err = server.keeper.eventService.EventManager(goCtx).EmitKV(
		goCtx,
		types.PoolCreatedEvent,
		event.Attribute{
			Key:   "pool_id",
			Value: strconv.FormatUint(pool.Id, 10),
		},
		event.Attribute{
			Key:   "pool_address",
			Value: pool.PoolAddress,
		},
		event.Attribute{
			Key:   "token_denom",
			Value: pool.TokenDenom,
		},
		event.Attribute{
			Key:   "pair_denom",
			Value: pool.PairDenom,
		},
		event.Attribute{
			Key:   "creator",
			Value: pool.Creator,
		},
		event.Attribute{
			Key:   "initial_token_denom",
			Value: tokenAmount.Denom,
		},
		event.Attribute{
			Key:   "initial_token_amount",
			Value: tokenAmount.Amount.String(),
		},
		event.Attribute{
			Key:   "initial_pair_denom",
			Value: pairAmount.Denom,
		},
		event.Attribute{
			Key:   "initial_pair_amount",
			Value: pairAmount.Amount.String(),
		},
		event.Attribute{
			Key:   "token_name",
			Value: pool.TokenName,
		},
		event.Attribute{
			Key:   "token_symbol",
			Value: pool.TokenSymbol,
		},
		event.Attribute{
			Key:   "description",
			Value: pool.TokenInformation.Description,
		},
		event.Attribute{
			Key:   "image",
			Value: pool.TokenInformation.Image,
		},
		event.Attribute{
			Key:   "website",
			Value: pool.TokenInformation.Website,
		},
		event.Attribute{
			Key:   "twitter",
			Value: pool.TokenInformation.Twitter,
		},
		event.Attribute{
			Key:   "telegram",
			Value: pool.TokenInformation.Telegram,
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreatePoolResponse{
		Id:                 strconv.FormatUint(pool.Id, 10),
		PoolAddress:        pool.PoolAddress,
		InitialTokenAmount: tokenAmount,
		InitialPairAmount:  pairAmount,
	}, nil
}

func (server msgServer) BuyToken(goCtx context.Context, msg *types.MsgBuyToken) (*types.MsgBuyTokenResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	pool, err := server.keeper.GetPool(goCtx, msg.Id)
	if err != nil {
		return nil, err
	}

	amountIn, amountOut, err := server.keeper.BuyToken(goCtx, msg.Id, sender, *msg.Amount)
	if err != nil {
		return nil, err
	}

	err = server.keeper.eventService.EventManager(goCtx).EmitKV(
		goCtx,
		types.BuyTokenEvent,
		event.Attribute{
			Key:   "pool_id",
			Value: strconv.FormatUint(msg.Id, 10),
		},
		event.Attribute{
			Key:   "pool_address",
			Value: pool.PoolAddress,
		},
		event.Attribute{
			Key:   "amount_in_denom",
			Value: amountIn.Denom,
		},
		event.Attribute{
			Key:   "amount_in_amount",
			Value: amountIn.Amount.String(),
		},
		event.Attribute{
			Key:   "amount_out_denom",
			Value: amountOut.Denom,
		},
		event.Attribute{
			Key:   "amount_out_amount",
			Value: amountOut.Amount.String(),
		},
		event.Attribute{
			Key:   "sender",
			Value: msg.FromAddress,
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgBuyTokenResponse{
		AmountIn:  amountIn,
		AmountOut: amountOut,
	}, nil
}

func (server msgServer) SellToken(goCtx context.Context, msg *types.MsgSellToken) (*types.MsgSellTokenResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	pool, err := server.keeper.GetPool(goCtx, msg.Id)
	if err != nil {
		return nil, err
	}

	amountIn, amountOut, err := server.keeper.SellToken(goCtx, msg.Id, sender, *msg.Amount)
	if err != nil {
		return nil, err
	}

	err = server.keeper.eventService.EventManager(goCtx).EmitKV(
		goCtx,
		types.SellTokenEvent,
		event.Attribute{
			Key:   "pool_id",
			Value: strconv.FormatUint(msg.Id, 10),
		},
		event.Attribute{
			Key:   "pool_address",
			Value: pool.PoolAddress,
		},
		event.Attribute{
			Key:   "amount_in_denom",
			Value: amountIn.Denom,
		},
		event.Attribute{
			Key:   "amount_in_amount",
			Value: amountIn.Amount.String(),
		},
		event.Attribute{
			Key:   "amount_out_denom",
			Value: amountOut.Denom,
		},
		event.Attribute{
			Key:   "amount_out_amount",
			Value: amountOut.Amount.String(),
		},
		event.Attribute{
			Key:   "sender",
			Value: msg.FromAddress,
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSellTokenResponse{
		AmountIn:  amountIn,
		AmountOut: amountOut,
	}, nil
}
