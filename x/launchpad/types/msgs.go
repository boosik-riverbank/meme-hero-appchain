package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewCreatePoolMsg(
	fromAddress sdk.AccAddress,
	tokenName,
	tokenSymbol,
	tokenDenom,
	pairDenom,
	description,
	imageUrl,
	website,
	telegram,
	twitter string,
	initialQuantity *sdk.Coin) *MsgCreatePool {
	return &MsgCreatePool{
		FromAddress:     fromAddress.String(),
		TokenName:       tokenName,
		TokenSymbol:     tokenSymbol,
		TokenDenom:      tokenDenom,
		PairDenom:       pairDenom,
		InitialQuantity: initialQuantity,
		Description:     description,
		Image:           imageUrl,
		Website:         website,
		Twitter:         twitter,
		Telegram:        telegram,
	}
}

func NewBuyTokenMsg(fromAddress sdk.AccAddress, id uint64, amount *sdk.Coin) *MsgBuyToken {
	return &MsgBuyToken{
		FromAddress: fromAddress.String(),
		Id:          id,
		Amount:      amount,
	}
}

func NewSellTokenMsg(fromAddress sdk.AccAddress, id uint64, amount *sdk.Coin) *MsgSellToken {
	return &MsgSellToken{
		FromAddress: fromAddress.String(),
		Id:          id,
		Amount:      amount,
	}
}
