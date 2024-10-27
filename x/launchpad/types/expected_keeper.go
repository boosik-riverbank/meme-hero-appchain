package types

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icacontroltypes "github.com/cosmos/interchain-security/v6/x/intertx/types"
)

type BankKeeper interface {
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type AuthKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
}

type InterTxKeeper interface {
	GetButler(ctx context.Context, id uint64) (*icacontroltypes.Butler, error)
	SendTx(ctx sdk.Context, owner, connectionId string, msgs []sdk.Msg, timeout uint64) error
	AddQueue(ctx sdk.Context, id uint64, targetChainId uint64, token, pair sdk.Coin) error
	TransferAssets(
		ctx context.Context,
		id uint64,
		sender sdk.AccAddress, receiver string,
		token sdk.Coin, memo string) error
	GetNextQueueId(ctx sdk.Context) (uint64, error)
	SetNextQueueId(ctx sdk.Context, id uint64) error
}
