package keeper

import (
	"math/big"
)

type BondingCurveConfig struct {
	InitMarketCap *big.Int
	EndMarketCap  *big.Int
	// SalePercent    *big.Int
	SalePercent    uint64
	BasisPoint     uint64
	TotalSupply    *big.Int
	PricePrecision uint8
}

var defaultBondingCurveConfig = BondingCurveConfig{
	InitMarketCap:  big.NewInt(30_000_000),
	EndMarketCap:   big.NewInt(120_000_000),
	SalePercent:    8000,
	BasisPoint:     10000,
	TotalSupply:    big.NewInt(1_073_000_191_000_000),
	PricePrecision: 6,
}

func (k Keeper) getK() *big.Int {
	return new(big.Int).Mul(defaultBondingCurveConfig.InitMarketCap, defaultBondingCurveConfig.TotalSupply)
}
