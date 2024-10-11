package keeper

import (
	"fmt"
	"math/big"
)

type BondingCurveConfig struct {
	InitMarketCap  *big.Int
	EndMarketCap   *big.Int
	SalePercent    *big.Int
	TotalSupply    *big.Int
	PricePrecision uint8
}

func (k Keeper) getDefaultBondingCurveConfig() (*BondingCurveConfig, error) {
	initMarketCap, ok := new(big.Int).SetString("30000000", 10)
	if !ok {
		return nil, fmt.Errorf("failed to set default bonding curve config")
	}

	endMarketCap, ok := new(big.Int).SetString("120000000", 10)
	if !ok {
		return nil, fmt.Errorf("failed to set default bonding curve config")
	}
	totalSupply, ok := new(big.Int).SetString("1000000000000000", 10)
	if !ok {
		return nil, fmt.Errorf("failed to set default bonding curve config")
	}

	bondingCurveConfig := BondingCurveConfig{
		InitMarketCap:  initMarketCap,
		EndMarketCap:   endMarketCap,
		SalePercent:    big.NewInt(8000),
		TotalSupply:    totalSupply,
		PricePrecision: 18,
	}

	return &bondingCurveConfig, nil
}

func (k Keeper) CalculateConfig() (*big.Int, *big.Int, *big.Int, error) {
	bondingCurveConfig, err := k.getDefaultBondingCurveConfig()
	if err != nil {
		return nil, nil, nil, err
	}

	//initMarketCap, err := k.calculateInitMarketCap(bondingCurveConfig.SalePercent, bondingCurveConfig.EndMarketCap)
	//if err != nil {
	//	return nil, nil, nil, err
	//}
	initMarketCap, ok := new(big.Int).SetString("30000000", 10)
	if !ok {
		return nil, nil, nil, fmt.Errorf("failed to set default bonding curve config")
	}

	saleAmount, err := MulDivRoundingUp(bondingCurveConfig.TotalSupply, bondingCurveConfig.SalePercent, big.NewInt(basisPoint))
	if err != nil {
		return nil, nil, nil, err
	}

	nativeOffset, err := k.calculateNativeOffset(bondingCurveConfig.EndMarketCap, initMarketCap)
	if err != nil {
		return nil, nil, nil, err
	}

	tokenOffset, err := k.CalculateTokenOffset(bondingCurveConfig.EndMarketCap, initMarketCap, bondingCurveConfig.SalePercent, bondingCurveConfig.TotalSupply)
	if err != nil {
		return nil, nil, nil, err
	}

	return saleAmount, tokenOffset, nativeOffset, nil
}
