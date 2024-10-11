package test_utils

import (
	"context"
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/types"
)

type MockBankKeeper struct {
	balances map[string]sdk.Coin
}

func NewMockBankKeeper() types.BankKeeper {
	return &MockBankKeeper{
		balances: make(map[string]sdk.Coin),
	}
}

func (m MockBankKeeper) GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	balance, ok := m.balances[addr.String()+denom]
	if !ok {
		return sdk.NewCoin(denom, math.NewInt(0))
	}

	return balance
}

func (m MockBankKeeper) SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error {
	for _, v := range amt {
		senderBalance, ok := m.balances[fromAddr.String()+v.Denom]
		if !ok {
			senderBalance = sdk.NewCoin(v.Denom, math.NewInt(0))
		}
		if senderBalance.Amount.LT(v.Amount) {
			return fmt.Errorf("insufficient balance")
		}

		receiverBalance, ok := m.balances[toAddr.String()+v.Denom]
		if !ok {
			receiverBalance = sdk.NewCoin(v.Denom, math.NewInt(0))
		}
		m.balances[fromAddr.String()+v.Denom] = sdk.NewCoin(v.Denom, senderBalance.Amount.Sub(v.Amount))
		m.balances[toAddr.String()+v.Denom] = sdk.NewCoin(v.Denom, receiverBalance.Amount.Add(v.Amount))
	}

	return nil
}

func (m MockBankKeeper) MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	acc := authtypes.NewEmptyModuleAccount(moduleName, authtypes.Minter)
	for _, v := range amt {
		exist, ok := m.balances[acc.GetAddress().String()+v.Denom]
		if !ok {
			exist = sdk.NewCoin(v.Denom, math.NewInt(0))
		}

		m.balances[acc.GetAddress().String()+v.Denom] = sdk.NewCoin(v.Denom, exist.Amount.Add(v.Amount))
	}

	return nil
}

func (m MockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	acc := authtypes.NewEmptyModuleAccount(senderModule, authtypes.Minter)
	return m.SendCoins(ctx, acc.GetAddress(), recipientAddr, amt)
}
