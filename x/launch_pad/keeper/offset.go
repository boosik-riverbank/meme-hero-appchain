package keeper

import (
	"fmt"
	"math/big"
)

func (k Keeper) calculateNativeOffset(endMarketCap *big.Int, initMarketCap *big.Int) (*big.Int, error) {
	res, err := MulDivRoundingUp(endMarketCap, initMarketCap, new(big.Int).Sub(endMarketCap, initMarketCap))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// TODO : temp use
const (
	basisPoint = 10000
	fee        = 100
)

func (k Keeper) CalculateTokenOffset(endMarketCap *big.Int, initMarketCap *big.Int, salePercent *big.Int, totalSupply *big.Int) (*big.Int, error) {
	if endMarketCap.Cmp(initMarketCap) <= 0 {
		return nil, fmt.Errorf("invalid market cap")
	}
	temp := new(big.Int).Mul(endMarketCap, new(big.Int).Sub(big.NewInt(basisPoint), salePercent))
	res, err := MulDivRoundingUp(
		totalSupply,
		new(big.Int).Add(new(big.Int).Mul(salePercent, initMarketCap), temp),
		new(big.Int).Mul(new(big.Int).Sub(endMarketCap, initMarketCap), big.NewInt(basisPoint)),
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (k Keeper) calculateInitMarketCap(salePercent *big.Int, endMarketCap *big.Int) (*big.Int, error) {
	// temp := (basisPoint - salePercent) * (basisPoint - salePercent)
	temp := new(big.Int).Mul(
		new(big.Int).Sub(big.NewInt(basisPoint), salePercent),
		new(big.Int).Sub(big.NewInt(basisPoint), salePercent),
	)
	res, err := MulDivRoundingUp(temp, endMarketCap, new(big.Int).Mul(salePercent, salePercent))
	if err != nil {
		return nil, err
	}
	return res, nil
}
