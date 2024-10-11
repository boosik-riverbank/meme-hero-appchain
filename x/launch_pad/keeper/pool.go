package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/types"
	"math/big"
)

var (
	// 0.01
	swapFee = big.NewInt(1_000_0)
	// 1
	createFee = big.NewInt(1_000_000)
)

func (k Keeper) quoteAmountIn2(amountOut *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	saleAmount, tokenOffset, nativeOffset, err := k.CalculateConfig()
	if amountOut.Cmp(saleAmount) == 1 {
		amountOut = saleAmount
	}

	reserveIn := nativeOffset
	reserveOut := new(big.Int).Add(tokenOffset, saleAmount)
	amountWithoutFee, err := MulDivRoundingUp(reserveIn, amountOut, new(big.Int).Sub(reserveOut, amountOut))
	if err != nil {
		return nil, nil, nil, err
	}

	nativeFee, err := MulDivRoundingUp(amountWithoutFee, swapFee, big.NewInt(basisPoint))
	if err != nil {
		return nil, nil, nil, err
	}

	return amountWithoutFee, amountOut, nativeFee, nil
}

func (k Keeper) quoteAmountOut2(amountIn *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	saleAmount, tokenOffset, nativeOffset, err := k.CalculateConfig()
	reserveIn := nativeOffset
	reserveOut := new(big.Int).Add(tokenOffset, saleAmount)
	nativeFee, err := MulDivRoundingUp(amountIn, swapFee, big.NewInt(basisPoint))
	if err != nil {
		return nil, nil, nil, err
	}

	amountOut, err := MulDiv(amountIn, reserveOut, new(big.Int).Add(reserveIn, amountIn))
	if err != nil {
		return nil, nil, nil, err
	}

	if amountOut.Cmp(saleAmount) == 1 {
		amountIn, _, _, err := k.quoteAmountIn2(amountOut)
		if err != nil {
			return nil, nil, nil, err
		}

		nativeFee, err = MulDivRoundingUp(amountIn, swapFee, big.NewInt(basisPoint))
		if err != nil {
			return nil, nil, nil, err
		}

		amountOut, err = MulDiv(amountIn, reserveOut, new(big.Int).Add(reserveIn, amountIn))
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return amountOut, amountIn, nativeFee, nil
}

func (k Keeper) CreateNewPool(ctx context.Context, msg *types.MsgCreatePool, poolAddress sdk.AccAddress) (uint64, error) {
	amount, ok := new(big.Int).SetString(msg.Amount, 10)
	if !ok {
		return 0, fmt.Errorf("invalid amount: %s", msg.Amount)
	}

	config := defaultBondingCurveConfig

	k.length++
	poolID := k.length

	pool := types.BondingCurvePool{
		Id:         poolID,
		TokenDenom: msg.TokenDenom,
		Creator:    msg.Creator,
		PairDenom:  msg.PairDenom,
	}

	err := k.pools.Set(ctx, poolID, pool)
	if err != nil {
		return 0, err
	}

	// Mint token to pool address
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(msg.TokenDenom, math.NewIntFromBigInt(config.TotalSupply))))
	if err != nil {
		return 0, err
	}

	// Send to created pool
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, poolAddress, sdk.NewCoins(sdk.NewCoin(msg.TokenDenom, math.NewIntFromBigInt(config.TotalSupply))))
	if err != nil {
		return 0, err
	}

	if math.NewIntFromBigInt(amount).GT(math.NewInt(0)) {
		if msg.IsNativeIn {
			creator, err := sdk.AccAddressFromBech32(msg.Creator)
			if err != nil {
				return 0, err
			}

			// createFee
			err = k.bankKeeper.SendCoins(ctx, creator, k.feeCollector, sdk.NewCoins(sdk.NewCoin(msg.Value.Denom, math.NewIntFromBigInt(createFee))))
			if err != nil {
				return 0, err
			}

			_, err = k.SwapExactIn(ctx, poolID, actualAmountIn, big.NewInt(1), true, creator, *msg.Value)
			if err != nil {
				return 0, err
			}

			err = k.bankKeeper.SendCoins(ctx, creator, poolAddress, sdk.NewCoins(sdk.NewCoin(msg.Value.Denom, math.NewIntFromBigInt(actualAmountIn))))
			if err != nil {
				return 0, err
			}
		} else {
			creator, err := sdk.AccAddressFromBech32(msg.Creator)
			if err != nil {
				return 0, err
			}

			err = k.bankKeeper.SendCoins(ctx, creator, poolAddress, sdk.NewCoins(sdk.NewCoin(msg.Value.Denom, math.NewIntFromBigInt(actualNativeIn))))
			if err != nil {
				return 0, err
			}

			// createFee
			err = k.bankKeeper.SendCoins(ctx, creator, k.feeCollector, sdk.NewCoins(sdk.NewCoin(msg.Value.Denom, math.NewIntFromBigInt(createFee))))
			if err != nil {
				return 0, err
			}

			_, err = k.SwapExactOut(ctx, poolID, amount, new(big.Int).Add(amount, nativeFee), true, creator, *msg.Value)
			if err != nil {
				return 0, err
			}
		}
	}

	return poolID, nil
}

func (k Keeper) DeletePool(ctx context.Context, poolID uint64) error {
	err := k.pools.Remove(ctx, poolID)
	if err != nil {
		return errorsmod.Wrapf(err, "failed to delete existing pool")
	}

	return nil
}

func (k Keeper) GetPools(ctx context.Context) ([]*types.BondingCurvePool, error) {
	res := make([]*types.BondingCurvePool, 0)
	err := k.pools.Walk(ctx, nil, func(key uint64, value types.BondingCurvePool) (stop bool, err error) {
		res = append(res, &value)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (k Keeper) GetPool(ctx context.Context, poolID uint64) (*types.BondingCurvePool, error) {
	pool, err := k.pools.Get(ctx, poolID)
	if err != nil {
		return nil, err
	}

	return &pool, nil
}

func (k Keeper) GetPoolAmount(ctx context.Context, poolID uint64, denom string) (*sdk.Coin, error) {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return nil, err
	}

	address, err := sdk.AccAddressFromBech32(pool.Address)
	if err != nil {
		return nil, err
	}

	res := k.bankKeeper.GetBalance(ctx, address, denom)
	return &res, nil
}

func (k Keeper) GetPoolParams(poolID uint64) {
	// TODO: implements it!
}
