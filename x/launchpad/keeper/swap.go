package keeper

import (
	"context"
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

func (k Keeper) BuyToken(ctx context.Context, poolID uint64, user sdk.AccAddress, wantToBuy sdk.Coin) (*sdk.Coin, *sdk.Coin, error) {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return nil, nil, err
	}

	if pool.IsListed {
		return nil, nil, fmt.Errorf("this pool is alreadt listed. you cannot use this pool")
	}

	if pool.PairDenom != wantToBuy.Denom {
		return nil, nil, fmt.Errorf("invalid pair token to buy")
	}

	tokenReserve, pairReserve := k.GetTokenPair(ctx, poolID)
	res, err := k.QuoteAmountOut(wantToBuy.Amount.BigInt(), tokenReserve.Amount.BigInt(), new(big.Int).Add(pairReserve.Amount.BigInt(), defaultBondingCurveConfig.InitMarketCap))
	if err != nil {
		return nil, nil, err
	}

	t := new(big.Int).Mul(new(big.Int).SetUint64(defaultBondingCurveConfig.BasisPoint-defaultBondingCurveConfig.SalePercent), defaultBondingCurveConfig.TotalSupply)
	limit := new(big.Int).Quo(t, new(big.Int).SetUint64(defaultBondingCurveConfig.BasisPoint))
	fmt.Printf("limit: %v\n", limit.String())
	if tokenReserve.Amount.Sub(math.NewIntFromBigInt(res)).LT(math.NewIntFromBigInt(limit)) {
		// Cannot sell because of limit
		return nil, nil, fmt.Errorf("cannot buy token because 80 percent of total supply is already sold")
	}

	poolAddress, err := sdk.AccAddressFromBech32(pool.PoolAddress)
	err = k.bankKeeper.SendCoins(ctx, poolAddress, user, sdk.Coins{sdk.NewCoin(pool.TokenDenom, math.NewIntFromBigInt(res))})
	if err != nil {
		return nil, nil, err
	}

	err = k.bankKeeper.SendCoins(ctx, user, poolAddress, sdk.Coins{sdk.NewCoin(pool.PairDenom, wantToBuy.Amount)})
	if err != nil {
		return nil, nil, err
	}

	currentMarketCap := k.GetCurrentMarketCap(ctx, poolID)
	fmt.Printf("currentMarketCap: %v\n", currentMarketCap.String())
	if math.NewIntFromBigInt(currentMarketCap).GT(math.NewInt(410_000000)) {
		// go to dex!!!
		fmt.Printf("TODO: Launch to DEX!!\n")
		err = k.SetListPool(ctx, poolID)
		if err != nil {
			return nil, nil, err
		}
	}

	return &sdk.Coin{Denom: wantToBuy.Denom, Amount: wantToBuy.Amount}, &sdk.Coin{Denom: pool.TokenDenom, Amount: math.NewIntFromBigInt(res)}, nil
}

func (k Keeper) QuoteAmountOut(wantToBuy, tokenReserve, pairReserve *big.Int) (*big.Int, error) {
	// d0 = k / (pairReserve - wantToBuy)
	d0, err := Div(k.getK(), new(big.Int).Add(pairReserve, wantToBuy))
	if err != nil {
		return nil, err
	}

	// d1 = tokenReserve - d0
	d1 := new(big.Int).Sub(tokenReserve, d0)
	return d1, nil
}

func (k Keeper) SellToken(ctx context.Context, poolID uint64, user sdk.AccAddress, wantToSell sdk.Coin) (*sdk.Coin, *sdk.Coin, error) {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return nil, nil, err
	}

	if pool.TokenDenom != wantToSell.Denom {
		return nil, nil, fmt.Errorf("invalid pair token to sell")
	}

	tokenReserve, pairReserve := k.GetTokenPair(ctx, poolID)

	res, err := k.QuoteAmountIn(wantToSell.Amount.BigInt(), tokenReserve.Amount.BigInt(), new(big.Int).Add(pairReserve.Amount.BigInt(), defaultBondingCurveConfig.InitMarketCap))
	if err != nil {
		return nil, nil, err
	}

	poolAddress, err := sdk.AccAddressFromBech32(pool.PoolAddress)
	// User -> Pool, send token
	err = k.bankKeeper.SendCoins(ctx, user, poolAddress, sdk.Coins{sdk.NewCoin(pool.TokenDenom, wantToSell.Amount)})
	if err != nil {
		return nil, nil, err
	}

	// Pool -> User, send pair token
	err = k.bankKeeper.SendCoins(ctx, poolAddress, user, sdk.Coins{sdk.NewCoin(pool.PairDenom, math.NewIntFromBigInt(res))})
	if err != nil {
		return nil, nil, err
	}

	return &sdk.Coin{Denom: pool.TokenDenom, Amount: wantToSell.Amount}, &sdk.Coin{Denom: pool.PairDenom, Amount: math.NewIntFromBigInt(res)}, nil
}

func (k Keeper) QuoteAmountIn(wantToSell, tokenReserve, pairReserve *big.Int) (*big.Int, error) {
	// d0 = nativeReserve + wantToSell
	d0, err := Div(k.getK(), new(big.Int).Add(tokenReserve, wantToSell))
	if err != nil {
		return nil, err
	}

	// d1 = tokenReserve - d0
	d1 := new(big.Int).Sub(pairReserve, d0)
	return d1, nil
}

func (k Keeper) GetCurrentPrice(ctx context.Context, poolID uint64) *big.Float {
	tokenReserve, pairReserve := k.GetTokenPair(ctx, poolID)
	if pairReserve.IsZero() {
		return big.NewFloat(0)
	}
	a := new(big.Float).Add(new(big.Float).SetInt(big.NewInt(30_000000)), new(big.Float).SetInt(pairReserve.Amount.BigInt()))
	b := new(big.Float).SetInt(tokenReserve.Amount.BigInt())
	price := new(big.Float).Quo(a, b)
	return price
}

func (k Keeper) GetCurrentMarketCap(ctx context.Context, poolID uint64) *big.Int {
	currentPrice := k.GetCurrentPrice(ctx, poolID)
	mc := new(big.Float).Mul(currentPrice, new(big.Float).SetInt(defaultBondingCurveConfig.TotalSupply))
	res, _ := mc.Int(new(big.Int))
	return res
}

func (k Keeper) GetTokenPair(ctx context.Context, poolID uint64) (tokenReserve *sdk.Coin, pairReserve *sdk.Coin) {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return nil, nil
	}
	if pool == nil {
		return nil, nil
	}

	poolAddress, err := sdk.AccAddressFromBech32(pool.PoolAddress)
	a := k.bankKeeper.GetBalance(ctx, poolAddress, pool.TokenDenom)
	b := k.bankKeeper.GetBalance(ctx, poolAddress, pool.PairDenom)
	tokenReserve = &a
	pairReserve = &b
	return tokenReserve, pairReserve
}
