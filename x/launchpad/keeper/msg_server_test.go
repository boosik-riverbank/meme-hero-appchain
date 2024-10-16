package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launchpad/types"
	"math/big"
	"strconv"
)

func (suite KeeperTestSuite) TestTransactionFlow() {
	// 1. Setup Test
	creatorAddr, _ := sdk.AccAddressFromBech32("meme1j49359c08v66rthaqutfh6wc7k04sssp843kqj")

	_ = suite.bankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("uatom", math.NewInt(10000_000000))))
	_ = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, creatorAddr, sdk.NewCoins(sdk.NewCoin("uatom", math.NewInt(1000_000000))))

	pool, err := suite.msgServer.CreatePool(suite.ctx, &types.MsgCreatePool{
		FromAddress: creatorAddr.String(),
		TokenName:   "dogecoin",
		TokenSymbol: "doge",
		TokenDenom:  "udoge",
		PairDenom:   "uatom",
		InitialQuantity: &sdk.Coin{
			Denom:  "uatom",
			Amount: math.NewInt(1_000000),
		},
		Description: "test doge coin",
		Image:       "https://image.com",
		Website:     "https://website.com",
		Telegram:    "https://telegram.com",
		Twitter:     "https://twitter.com",
	})
	suite.Require().NoError(err)
	fmt.Printf("\n##### Created Pool and user buy 1 ATOM intially, id=%v, poolAddress=%v #####\n", pool.Id, pool.PoolAddress)

	fmt.Printf(">>> Check user and pool balance <<<\n")
	id, err := strconv.Atoi(pool.Id)
	tokenReserve, pairReserve := suite.keeper.GetTokenPair(suite.ctx, uint64(id))
	userPairBalance := suite.bankKeeper.GetBalance(suite.ctx, creatorAddr, "uatom")
	userTokenBalance := suite.bankKeeper.GetBalance(suite.ctx, creatorAddr, "udoge")
	fmt.Printf("Pool) udoge=%v, uatom=%v\n", tokenToFloat(*tokenReserve.Amount.BigInt()), tokenToFloat(*pairReserve.Amount.BigInt()))
	fmt.Printf("User) udoge=%v, uatom=%v\n", tokenToFloat(*userTokenBalance.Amount.BigInt()), tokenToFloat(*userPairBalance.Amount.BigInt()))
	tokenReserve, pairReserve = suite.keeper.GetTokenPair(suite.ctx, uint64(id))
	userPairBalance = suite.bankKeeper.GetBalance(suite.ctx, creatorAddr, "uatom")
	userTokenBalance = suite.bankKeeper.GetBalance(suite.ctx, creatorAddr, "udoge")
	price := suite.keeper.GetCurrentPrice(suite.ctx, uint64(id))
	mc := suite.keeper.GetCurrentMarketCap(suite.ctx, uint64(id))
	fmt.Printf("Pool) udoge=%v, uatom=%v\n", tokenToFloat(*tokenReserve.Amount.BigInt()), tokenToFloat(*pairReserve.Amount.BigInt()))
	fmt.Printf("User) udoge=%v, uatom=%v\n", tokenToFloat(*userTokenBalance.Amount.BigInt()), tokenToFloat(*userPairBalance.Amount.BigInt()))
	fmt.Printf("Current Price=%vuatom\n", price.String())
	fmt.Printf("Current MarketCap=%vuatom\n", mc.String())

	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)
	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)
	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)
	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)
	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)
	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)
	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)

	err = suite.TestSell(creatorAddr, 0, math.NewInt(10000000_000000))
	suite.Require().NoError(err)
	err = suite.TestSell(creatorAddr, 0, math.NewInt(10000000_000000))
	suite.Require().NoError(err)
	err = suite.TestSell(creatorAddr, 0, math.NewInt(10000000_000000))
	suite.Require().NoError(err)
	err = suite.TestSell(creatorAddr, 0, math.NewInt(10000000_000000))
	suite.Require().NoError(err)
	err = suite.TestSell(creatorAddr, 0, math.NewInt(10000000_000000))
	suite.Require().NoError(err)

	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)
	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)
	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().NoError(err)

	// Listed
	err = suite.TestBuy(creatorAddr, 0, math.NewInt(10_000000))
	suite.Require().Error(err)

}

func (suite KeeperTestSuite) TestBuy(creatorAddr sdk.AccAddress, poolID uint64, wantToBuy math.Int) error {
	fmt.Printf("\n##### Buy token with %v ATOM #####\n", wantToBuy.Quo(math.NewInt(1_000000)))
	res, err := suite.msgServer.BuyToken(suite.ctx, &types.MsgBuyToken{
		Id:          poolID,
		FromAddress: creatorAddr.String(),
		Amount: &sdk.Coin{
			Denom:  "uatom",
			Amount: wantToBuy,
		},
	})
	if err != nil {
		return err
	}
	fmt.Printf("User buy: %v udoge\n", res.AmountOut)
	fmt.Printf("User pay: %v uatom\n", res.AmountIn)
	fmt.Printf(">>> Check user and pool balance <<<\n")
	tokenReserve, pairReserve := suite.keeper.GetTokenPair(suite.ctx, poolID)
	userPairBalance := suite.bankKeeper.GetBalance(suite.ctx, creatorAddr, "uatom")
	userTokenBalance := suite.bankKeeper.GetBalance(suite.ctx, creatorAddr, "udoge")
	mc := suite.keeper.GetCurrentMarketCap(suite.ctx, poolID)
	price := suite.keeper.GetCurrentPrice(suite.ctx, poolID)
	fmt.Printf("Pool) udoge=%v, uatom=%v\n", tokenToFloat(*tokenReserve.Amount.BigInt()), tokenToFloat(*pairReserve.Amount.BigInt()))
	fmt.Printf("User) udoge=%v, uatom=%v\n", tokenToFloat(*userTokenBalance.Amount.BigInt()), tokenToFloat(*userPairBalance.Amount.BigInt()))
	fmt.Printf("Current Price=%vuatom\n", price.String())
	fmt.Printf("Current MarketCap=%vuatom\n", mc.String())

	return nil
}

func (suite KeeperTestSuite) TestSell(creatorAddr sdk.AccAddress, poolID uint64, wantToSell math.Int) error {
	fmt.Printf("\n##### Sell token with %v DOGE #####\n", wantToSell.Quo(math.NewInt(1_000000)))
	res, err := suite.msgServer.SellToken(suite.ctx, &types.MsgSellToken{
		Id:          poolID,
		FromAddress: creatorAddr.String(),
		Amount: &sdk.Coin{
			Denom:  "udoge",
			Amount: wantToSell,
		},
	})
	if err != nil {
		return err
	}
	fmt.Printf("User sell: %v udoge\n", res.AmountIn)
	fmt.Printf("User get: %v uatom\n", res.AmountOut)
	fmt.Printf(">>> Check user and pool balance <<<\n")
	tokenReserve, pairReserve := suite.keeper.GetTokenPair(suite.ctx, poolID)
	userPairBalance := suite.bankKeeper.GetBalance(suite.ctx, creatorAddr, "uatom")
	userTokenBalance := suite.bankKeeper.GetBalance(suite.ctx, creatorAddr, "udoge")
	mc := suite.keeper.GetCurrentMarketCap(suite.ctx, poolID)
	price := suite.keeper.GetCurrentPrice(suite.ctx, poolID)
	fmt.Printf("Pool) udoge=%v, uatom=%v\n", tokenToFloat(*tokenReserve.Amount.BigInt()), tokenToFloat(*pairReserve.Amount.BigInt()))
	fmt.Printf("User) udoge=%v, uatom=%v\n", tokenToFloat(*userTokenBalance.Amount.BigInt()), tokenToFloat(*userPairBalance.Amount.BigInt()))
	fmt.Printf("Current Price=%vuatom\n", price.String())
	fmt.Printf("Current MarketCap=%vuatom\n", mc.String())
	return nil
}

func tokenToFloat(token big.Int) *big.Int {
	return &token
}
