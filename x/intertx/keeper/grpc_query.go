package keeper

import (
	"context"
	"github.com/cosmos/interchain-security/v6/x/intertx/types"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Keeper{}

func (k Keeper) Butlers(ctx context.Context, butlers *types.QueryButlers) (*types.QueryButlersResponse, error) {
	var res []types.Butler
	err := k.butlers.Walk(ctx, nil, func(key uint64, value types.Butler) (stop bool, err error) {
		res = append(res, value)
		return false, nil
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryButlersResponse{
		Butlers: res,
	}, nil
}

func (k Keeper) Butler(ctx context.Context, butler *types.QueryButler) (*types.QueryButlerResponse, error) {
	res, err := k.GetButler(ctx, butler.Id)
	if err != nil {
		return nil, err
	}
	return &types.QueryButlerResponse{
		Butler: *res,
	}, nil
}
