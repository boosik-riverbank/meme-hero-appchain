package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"math/big"
)

func GenerateNewAddress(moduleName string, prefix string, identifier []byte) sdk.AccAddress {
	key := append([]byte(prefix), identifier...)
	return address.Module(moduleName, key)
}

func MulDivRoundingUp(a, b, denominator *big.Int) (*big.Int, error) {
	if denominator.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("denominator is zero")
	}
	prod := new(big.Int).Mul(a, b)
	div := new(big.Float).Quo(new(big.Float).SetInt(prod), new(big.Float).SetInt(denominator))
	add := new(big.Float).Add(div, big.NewFloat(0.5))
	res, _ := add.Int(new(big.Int))
	return res, nil
}

func MulDiv(a, b, denominator *big.Int) (*big.Int, error) {
	if denominator.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("denominator is zero")
	}
	prod := new(big.Int).Mul(a, b)
	div := new(big.Float).Quo(new(big.Float).SetInt(prod), new(big.Float).SetInt(denominator))
	res, _ := div.Int(new(big.Int))
	return res, nil
}
