package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launchpad/types"
	"math/big"
)

var (
	// 0.01
	swapFee = big.NewInt(1_000_0)
	// 1
	createFee = big.NewInt(1_000_000)
)

func (k Keeper) CreateNewPool(ctx context.Context, msg *types.MsgCreatePool) (*types.BondingCurvePool, error) {
	config := defaultBondingCurveConfig

	set, err := k.count.Has(ctx)
	if err != nil {
		return nil, err
	}
	if !set {
		err = k.count.Set(ctx, uint64(0))
		if err != nil {
			return nil, err
		}
	}

	poolID, err := k.count.Get(ctx)
	if err != nil {
		return nil, err
	}

	poolAddress := GenerateNewAddress(types.ModuleName, "pool", sdk.Uint64ToBigEndian(poolID))
	pool := types.BondingCurvePool{
		Id:          poolID,
		TokenName:   msg.TokenName,
		TokenSymbol: msg.TokenSymbol,
		TokenDenom:  msg.TokenDenom,
		Creator:     msg.FromAddress,
		PairDenom:   msg.PairDenom,
		PoolAddress: poolAddress.String(),
		TokenInformation: &types.TokenInformation{
			Description: msg.Description,
			Image:       msg.Image,
			Website:     msg.Website,
			Telegram:    msg.Telegram,
			Twitter:     msg.Twitter,
		},
	}

	err = k.pools.Set(ctx, poolID, pool)
	if err != nil {
		return nil, err
	}

	// Mint token to pool address
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(msg.TokenDenom, math.NewIntFromBigInt(config.TotalSupply))))
	if err != nil {
		return nil, err
	}

	// Send to created pool
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, poolAddress, sdk.NewCoins(sdk.NewCoin(msg.TokenDenom, math.NewIntFromBigInt(config.TotalSupply))))
	if err != nil {
		return nil, err
	}

	creator, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	if msg.InitialQuantity.Amount.GT(math.NewInt(0)) {
		if msg.PairDenom != msg.InitialQuantity.Denom {
			return nil, fmt.Errorf("invalid coin denom of intial quantity")
		}
		initialBought, err := k.QuoteAmountOut(msg.InitialQuantity.Amount.BigInt(), config.TotalSupply, config.InitMarketCap)
		if err != nil {
			return nil, err
		}

		err = k.bankKeeper.SendCoins(ctx, poolAddress, creator, sdk.Coins{sdk.NewCoin(msg.TokenDenom, math.NewIntFromBigInt(initialBought))})
		if err != nil {
			return nil, err
		}

		err = k.bankKeeper.SendCoins(ctx, creator, poolAddress, sdk.Coins{sdk.NewCoin(msg.PairDenom, msg.InitialQuantity.Amount)})
		if err != nil {
			return nil, err
		}
	}

	err = k.count.Set(ctx, poolID+1)
	if err != nil {
		return nil, err
	}

	return &pool, nil
}

func (k Keeper) DeletePool(ctx context.Context, poolID uint64) error {
	err := k.pools.Remove(ctx, poolID)
	if err != nil {
		return errorsmod.Wrapf(err, "failed to delete existing pool")
	}

	return nil
}

func (k Keeper) GetPools(ctx context.Context) ([]*types.BondingCurvePool, error) {
	var res []*types.BondingCurvePool
	err := k.pools.Walk(ctx, nil, func(key uint64, value types.BondingCurvePool) (stop bool, err error) {
		k.logger.Info(fmt.Sprintf("walk: %d", key))
		res = append(res, &value)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (k Keeper) SetPool(ctx context.Context, poolID uint64, pool types.BondingCurvePool) error {
	err := k.pools.Set(ctx, poolID, pool)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetPool(ctx context.Context, poolID uint64) (*types.BondingCurvePool, error) {
	pool, err := k.pools.Get(ctx, poolID)
	if err != nil {
		return nil, err
	}

	return &pool, nil
}

func (k Keeper) GetPoolAmount(ctx context.Context, poolID uint64) (*sdk.Coin, *sdk.Coin, error) {
	tokenReserve, pairReserve := k.GetTokenPair(ctx, poolID)
	if tokenReserve == nil || pairReserve == nil {
		return nil, nil, fmt.Errorf("no token pair")
	}

	return tokenReserve, pairReserve, nil
}

func (k Keeper) SetListPool(ctx context.Context, poolID uint64) error {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return err
	}

	pool.IsListed = true
	err = k.SetPool(ctx, poolID, *pool)
	if err != nil {
		return err
	}

	return nil
}
