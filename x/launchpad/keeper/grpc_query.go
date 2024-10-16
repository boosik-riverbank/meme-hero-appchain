package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/interchain-security/v6/x/launchpad/types"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Keeper{}

func (k Keeper) Pools(ctx context.Context, pools *types.QueryPools) (*types.QueryPoolsResponse, error) {
	res, err := k.GetPools(ctx)
	if err != nil {
		return nil, fmt.Errorf("get count err: %v", err.Error())
	}

	var arr []types.BondingCurvePool
	for _, pool := range res {
		arr = append(arr, *pool)
	}

	set, err := k.count.Has(ctx)
	c := uint64(0)
	if set {
		c, err = k.count.Get(ctx)
		if err != nil {
			return nil, fmt.Errorf("get count err: %v", err.Error())
		}
	}

	return &types.QueryPoolsResponse{Count: c, Pools: arr}, nil
}

func (k Keeper) Pool(ctx context.Context, pool *types.QueryPool) (*types.QueryPoolResponse, error) {
	res, err := k.GetPool(ctx, pool.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryPoolResponse{Pool: *res}, nil
}

func (k Keeper) PoolAmount(ctx context.Context, query *types.QueryPoolAmount) (*types.QueryPoolAmountResponse, error) {
	tokenReserve, pairReserve, err := k.GetPoolAmount(ctx, query.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryPoolAmountResponse{
		Token: *tokenReserve,
		Pair:  *pairReserve,
	}, nil
}
