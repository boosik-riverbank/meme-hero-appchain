package keeper

import (
	"context"
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

func (k Keeper) SwapExactIn(
	ctx context.Context,
	poolID uint64,
	amountIn *big.Int,
	minimumReceive *big.Int,
	isBuyToken bool,
	user sdk.AccAddress,
	value sdk.Coin) (*big.Int, error) {
	// isBuyToken : amountIn -> amountNative
	// !isBuyToken : amountIn -> amountToken
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return nil, err
	}

	amountToken := new(big.Int).SetInt64(0)
	amountNative := new(big.Int).SetInt64(0)
	poolAddress, err := sdk.AccAddressFromBech32(pool.Address)

	tokenReserve, nativeReserve := k.getTokenPair(ctx, poolID)
	tokenOffset, ok := new(big.Int).SetString(pool.TokenOffset, 10)
	if !ok {
		return nil, fmt.Errorf("invalid token offset")
	}
	nativeOffset, ok := new(big.Int).SetString(pool.NativeOffset, 10)
	if !ok {
		return nil, fmt.Errorf("invalid native token offset")
	}
	fee, ok := new(big.Int).SetString(pool.Fee, 10)
	if !ok {
		return nil, fmt.Errorf("invalid fee")
	}

	basisPointBigInt := big.NewInt(basisPoint)

	if isBuyToken {
		amountOut, actualAmountIn, nativeFee, err := k.QuoteAmountOut(
			amountIn,
			tokenOffset,
			nativeOffset,
			tokenReserve.Amount.BigInt(),
			nativeReserve.Amount.BigInt(),
			isBuyToken,
			fee,
			basisPointBigInt,
		)
		if err != nil {
			return nil, err
		}
		if actualAmountIn.Cmp(amountIn) == -1 {
			amountIn = actualAmountIn
		}
		amountToken = amountOut
		amountNative = amountIn

		//if value.Amount.BigInt().Cmp(new(big.Int).Add(amountNative, nativeFee)) == 1 {
		//	charge1 := new(big.Int).Sub(value.Amount.BigInt(), amountNative)
		//	charge2 := new(big.Int).Sub(charge1, nativeFee)
		//	err = k.bankKeeper.SendCoins(ctx, poolAddress, user, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(charge2))))
		//	if err != nil {
		//		return nil, err
		//	}
		//}

		// Pool -> User, send meme token
		err = k.bankKeeper.SendCoins(ctx, poolAddress, user, sdk.NewCoins(sdk.NewCoin(pool.TokenDenom, math.NewIntFromBigInt(amountToken))))
		if err != nil {
			return nil, err
		}

		// User -> Pool, send native token
		err = k.bankKeeper.SendCoins(ctx, user, poolAddress, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(amountNative))))
		if err != nil {
			return nil, err
		}

		// User -> feeCollector, send fee
		err = k.bankKeeper.SendCoins(ctx, user, k.feeCollector, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(nativeFee))))
		if err != nil {
			return nil, err
		}

		// TODO : emit event

		// check can inject to DEX
		//updatedBalance := k.bankKeeper.GetBalance(ctx, poolAddress, pool.Address)
		//if updatedBalance.Amount.Uint64() >= pool.TokenOffset {
		//	// TODO : inject liquidity to DEX
		//}

		return amountToken, nil
	} else {
		amountOut, actualAmountIn, nativeFee, err := k.QuoteAmountOut(
			amountIn,
			tokenOffset,
			nativeOffset,
			tokenReserve.Amount.BigInt(),
			nativeReserve.Amount.BigInt(),
			isBuyToken,
			fee,
			basisPointBigInt,
		)
		if err != nil {
			return nil, err
		}
		if value.Amount.BigInt().Cmp(nativeFee) == -1 {
			return nil, fmt.Errorf("insufficient input")
		}
		if amountOut.Cmp(minimumReceive) == -1 {
			return nil, fmt.Errorf("insufficient output")
		}
		if actualAmountIn.Cmp(amountIn) == -1 {
			amountIn = actualAmountIn
		}
		amountToken, amountNative = amountIn, amountOut

		// Pool -> User, send native token
		err = k.bankKeeper.SendCoins(ctx, poolAddress, user, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(amountNative))))
		if err != nil {
			return nil, err
		}

		// User -> Pool, send meme token
		err = k.bankKeeper.SendCoins(ctx, user, poolAddress, sdk.NewCoins(sdk.NewCoin(pool.TokenDenom, math.NewIntFromBigInt(amountToken))))
		if err != nil {
			return nil, err
		}

		// User -> feeCollector, send fee
		err = k.bankKeeper.SendCoins(ctx, poolAddress, k.feeCollector, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(nativeFee))))
		if err != nil {
			return nil, err
		}

		// TODO : emit event

		// check can inject to DEX
		//updatedBalance := k.bankKeeper.GetBalance(ctx, poolAddress, pool.Address)
		//if updatedBalance.Amount.Uint64() >= pool.TokenOffset {
		//	// TODO : inject liquidity to DEX
		//}

		return amountNative, nil
	}
}

func (k Keeper) QuoteAmountIn(
	amountOut *big.Int,
	tokenOffset *big.Int,
	nativeOffset *big.Int,
	tokenReserve *big.Int,
	nativeReserve *big.Int,
	isBuyToken bool,
	fee *big.Int,
	basisPoint *big.Int) (amountIn *big.Int, actualAmountOut *big.Int, nativeFee *big.Int, err error) {
	actualAmountOut = new(big.Int).Set(amountOut)
	if isBuyToken {
		// amountIn -> amountNative
		if tokenReserve.Cmp(new(big.Int).Add(amountOut, tokenOffset)) == -1 {
			actualAmountOut = new(big.Int).Sub(nativeReserve, nativeOffset)
		}

		amountIn, nativeFee, err = k.quoteAmountIn(actualAmountOut, nativeReserve, tokenReserve, isBuyToken, fee, basisPoint)
		if err != nil {
			return nil, nil, nil, err
		}
	} else {
		// amountNative -> amountIn
		if nativeReserve.Cmp(new(big.Int).Add(amountOut, nativeOffset)) == -1 {
			actualAmountOut = new(big.Int).Sub(nativeReserve, nativeOffset)
		}

		amountIn, nativeFee, err = k.quoteAmountIn(actualAmountOut, tokenReserve, nativeReserve, isBuyToken, fee, basisPoint)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	if amountIn.Cmp(new(big.Int).SetInt64(0)) == 0 || nativeFee.Cmp(new(big.Int).SetInt64(0)) == 0 {
		return nil, nil, nil, fmt.Errorf("unexpected amount. amountIn=%v, nativeFee=%v", amountIn, nativeFee)
	}

	return amountIn, actualAmountOut, nativeFee, nil
}

func (k Keeper) quoteAmountIn(
	amountOut *big.Int,
	reserveIn *big.Int,
	reserveOut *big.Int,
	isNativeIn bool,
	fee *big.Int,
	basisPoint *big.Int) (*big.Int, *big.Int, error) {
	if isNativeIn {
		sub := new(big.Int).Sub(reserveOut, amountOut)
		if reserveIn.Cmp(big.NewInt(0)) == 0 {
			reserveIn = big.NewInt(10000)
		}
		amountInWithoutFee, err := MulDivRoundingUp(reserveIn, amountOut, sub)
		if err != nil {
			return nil, nil, err
		}

		nativeFee, err := MulDivRoundingUp(amountInWithoutFee, fee, basisPoint)
		if err != nil {
			return nil, nil, err
		}

		return amountInWithoutFee, nativeFee, nil
	} else {
		nativeFee, err := MulDivRoundingUp(amountOut, fee, basisPoint)
		if err != nil {
			return nil, nil, err
		}

		sub := new(big.Int).Sub(reserveOut, amountOut)
		amountInAfterFee, err := MulDivRoundingUp(amountOut, reserveIn, sub)
		if err != nil {
			return nil, nil, err
		}

		return amountInAfterFee, nativeFee, nil
	}
}

func (k Keeper) SwapExactOut(
	ctx context.Context,
	poolID uint64,
	amountOut *big.Int,
	maximumPay *big.Int,
	isBuyToken bool,
	user sdk.AccAddress,
	value sdk.Coin) (*big.Int, error) {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return nil, err
	}

	amountToken := new(big.Int).SetInt64(0)
	amountNative := new(big.Int).SetInt64(0)
	poolAddress, err := sdk.AccAddressFromBech32(pool.Address)

	tokenReserve, nativeReserve := k.getTokenPair(ctx, poolID)
	tokenOffset, ok := new(big.Int).SetString(pool.TokenOffset, 10)
	if !ok {
		return nil, fmt.Errorf("invalid token offset")
	}
	nativeOffset, ok := new(big.Int).SetString(pool.NativeOffset, 10)
	if !ok {
		return nil, fmt.Errorf("invalid native token offset")
	}
	fee := big.NewInt(fee)
	basisPointBigInt := big.NewInt(basisPoint)

	if isBuyToken {
		amountIn, actualAmountOut, nativeFee, err := k.QuoteAmountIn(
			amountOut,
			tokenOffset,
			nativeOffset,
			tokenReserve.Amount.BigInt(),
			nativeReserve.Amount.BigInt(),
			isBuyToken,
			fee,
			basisPointBigInt)
		if err != nil {
			return nil, err
		}
		if actualAmountOut.Cmp(amountOut) == -1 {
			amountOut = actualAmountOut
		}
		amountToken = amountOut
		amountNative = amountIn
		if value.Amount.BigInt().Cmp(new(big.Int).Add(amountNative, nativeFee)) == 1 {
			// value.Amount.Uint64()-amountNative-nativeFee
			charge1 := new(big.Int).Sub(value.Amount.BigInt(), amountNative)
			charge2 := new(big.Int).Sub(charge1, nativeFee)
			err = k.bankKeeper.SendCoins(ctx, poolAddress, user, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(charge2))))
			if err != nil {
				return nil, err
			}
		}

		// Pool -> User, send meme token
		err = k.bankKeeper.SendCoins(ctx, poolAddress, user, sdk.NewCoins(sdk.NewCoin(pool.TokenDenom, math.NewIntFromBigInt(amountToken))))
		if err != nil {
			return nil, err
		}

		// User -> Pool, send native token
		err = k.bankKeeper.SendCoins(ctx, user, poolAddress, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(amountNative))))
		if err != nil {
			return nil, err
		}

		// User -> feeCollector, send fee
		err = k.bankKeeper.SendCoins(ctx, poolAddress, k.feeCollector, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(nativeFee))))
		if err != nil {
			return nil, err
		}

		// TODO : emit event

		// check can inject to DEX
		// updatedBalance := k.bankKeeper.GetBalance(ctx, poolAddress, pool.Address)
		//if updatedBalance.Amount.Uint64() >= pool.TokenOffset {
		//	// TODO : inject liquidity to DEX
		//}

		return amountNative, nil
	} else {
		amountIn, actualAmountOut, nativeFee, err := k.QuoteAmountIn(
			amountOut,
			tokenOffset,
			nativeOffset,
			tokenReserve.Amount.BigInt(),
			nativeReserve.Amount.BigInt(),
			isBuyToken,
			fee,
			basisPointBigInt)
		if err != nil {
			return nil, err
		}
		if actualAmountOut.Cmp(amountOut) == -1 {
			amountOut = actualAmountOut
		}

		amountToken = amountIn
		amountNative = amountOut
		if amountToken.Cmp(maximumPay) == -1 {
			return nil, fmt.Errorf("insufficient input")
		}

		// Pool -> User, send native token
		err = k.bankKeeper.SendCoins(ctx, poolAddress, user, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(amountNative))))
		if err != nil {
			return nil, err
		}

		// User -> Pool, send meme token
		err = k.bankKeeper.SendCoins(ctx, user, poolAddress, sdk.NewCoins(sdk.NewCoin(pool.TokenDenom, math.NewIntFromBigInt(amountToken))))
		if err != nil {
			return nil, err
		}

		// User -> feeCollector, send fee
		err = k.bankKeeper.SendCoins(ctx, poolAddress, k.feeCollector, sdk.NewCoins(sdk.NewCoin(pool.NativeDenom, math.NewIntFromBigInt(nativeFee))))
		if err != nil {
			return nil, err
		}

		// TODO : emit event

		// check can inject to DEX
		updatedBalance := k.bankKeeper.GetBalance(ctx, poolAddress, pool.Address)
		if updatedBalance.Amount.BigInt().Cmp(tokenOffset) >= 0 {
			// TODO : inject liquidity to DEX
		}

		return amountToken, nil
	}
}

func (k Keeper) QuoteAmountOut(
	amountIn *big.Int,
	tokenOffset *big.Int,
	nativeOffset *big.Int,
	tokenReserve *big.Int,
	nativeReserve *big.Int,
	isBuyToken bool,
	fee *big.Int,
	basisPoint *big.Int) (amountOut *big.Int, actualAmountIn *big.Int, nativeFee *big.Int, err error) {
	actualAmountIn = amountIn

	if isBuyToken {
		amountOut, nativeFee, err = k.quoteAmountOut(amountIn, nativeReserve, tokenReserve, isBuyToken, fee, basisPoint)
		if err != nil {
			return nil, nil, nil, err
		}

		if tokenReserve.Cmp(new(big.Int).Add(amountOut, tokenOffset)) == -1 {
			amountOut = new(big.Int).Sub(tokenReserve, tokenOffset)
			actualAmountIn, nativeFee, err = k.quoteAmountIn(amountOut, nativeReserve, tokenReserve, isBuyToken, fee, basisPoint)
			if err != nil {
				return nil, nil, nil, err
			}
		}

		if amountOut.Cmp(new(big.Int).SetInt64(0)) == 0 || nativeFee.Cmp(new(big.Int).SetInt64(0)) == 0 {
			return nil, nil, nil, fmt.Errorf("unexpected amount. amountIn=%v, nativeFee=%v", amountIn, nativeFee)
		}

	} else {
		amountOut, nativeFee, err = k.quoteAmountOut(amountIn, tokenReserve, nativeReserve, isBuyToken, fee, basisPoint)
		if err != nil {
			return nil, nil, nil, err
		}

		if nativeReserve.Cmp(new(big.Int).Add(amountOut, nativeOffset)) == -1 {
			amountOut = new(big.Int).Sub(nativeReserve, nativeOffset)
			actualAmountIn, nativeFee, err = k.quoteAmountIn(amountOut, tokenReserve, nativeReserve, isBuyToken, fee, basisPoint)
		}

		if amountOut.Cmp(new(big.Int).SetInt64(0)) == 0 || nativeFee.Cmp(new(big.Int).SetInt64(0)) == 0 {
			return nil, nil, nil, fmt.Errorf("unexpected amount. amountIn=%v, nativeFee=%v", amountIn, nativeFee)
		}

	}
	return amountOut, actualAmountIn, nativeFee, nil

}

func (k Keeper) quoteAmountOut(
	amountIn *big.Int,
	reserveIn *big.Int,
	reserveOut *big.Int,
	isNativeIn bool,
	fee *big.Int,
	basisPoint *big.Int) (*big.Int, *big.Int, error) {
	if isNativeIn {
		nativeFee, err := MulDivRoundingUp(amountIn, fee, basisPoint)
		if err != nil {
			return nil, nil, err
		}

		amountOut, err := MulDiv(amountIn, reserveOut, new(big.Int).Add(reserveIn, amountIn))
		if err != nil {
			return nil, nil, err
		}

		return amountOut, nativeFee, nil
	} else {
		amountOutWithoutFee, err := MulDiv(amountIn, reserveOut, new(big.Int).Add(reserveIn, amountIn))
		if err != nil {
			return nil, nil, err
		}

		nativeFee, err := MulDivRoundingUp(amountOutWithoutFee, fee, basisPoint)
		if err != nil {
			return nil, nil, err
		}

		return amountOutWithoutFee, nativeFee, nil
	}
}

func (k Keeper) getTokenPair(ctx context.Context, poolID uint64) (tokenReserve sdk.Coin, atomReserve sdk.Coin) {
	pool, err := k.GetPool(ctx, poolID)
	if err != nil {
		return
	}
	if pool == nil {
		return
	}

	poolAddress, err := sdk.AccAddressFromBech32(pool.Address)
	tokenReserve = k.bankKeeper.GetBalance(ctx, poolAddress, pool.TokenDenom)
	atomReserve = k.bankKeeper.GetBalance(ctx, poolAddress, pool.NativeDenom)
	return tokenReserve, atomReserve
}
