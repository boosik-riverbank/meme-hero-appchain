package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launchpad/types"
)

func (suite *KeeperTestSuite) TestCreateNewPool() {
	// setup test
	require := suite.Require()
	creatorAddr, _ := sdk.AccAddressFromBech32("meme1j49359c08v66rthaqutfh6wc7k04sssp843kqj")

	_ = suite.bankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("uatom", math.NewInt(10000_000000))))
	_ = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, creatorAddr, sdk.NewCoins(sdk.NewCoin("uatom", math.NewInt(100_000000))))

	pool, err := suite.keeper.CreateNewPool(suite.ctx, &types.MsgCreatePool{
		TokenDenom:      "udoge",
		FromAddress:     creatorAddr.String(),
		PairDenom:       "uatom",
		InitialQuantity: &sdk.Coin{Denom: "uatom", Amount: math.NewInt(2_000000)},
	})
	require.NoError(err)
	require.Equal(pool.Id, uint64(0))

	result, err := suite.keeper.GetPool(suite.ctx, uint64(0))
	require.NoError(err)
	fmt.Printf("poolID: %v\n", result)

	tokenReserve, pairReserve := suite.keeper.GetTokenPair(suite.ctx, result.Id)
	fmt.Printf("tokenReserve=%v, pairReserve=%v\n", tokenReserve.String(), pairReserve.String())
}

func (suite *KeeperTestSuite) TestDeletePool() {}

func (suite *KeeperTestSuite) TestGetPools() {}

func (suite *KeeperTestSuite) TestGetPool() {}

func (suite *KeeperTestSuite) TestGetPoolAmount() {}
