package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/types"
	"github.com/stretchr/testify/require"
	"math/big"
)

func (suite *KeeperTestSuite) TestSwapExactIn() {
	// prepare to test
	suite.mockGetModuleAddress(types.ModuleName, launchPadAcc.GetAddress())

	fmt.Printf("Mint 1,000,000,000 DOGE!")
	creator, err := sdk.AccAddressFromBech32("meme1j49359c08v66rthaqutfh6wc7k04sssp843kqj")
	suite.bankKeeper.MintCoins(
		suite.ctx,
		types.ModuleName,
		sdk.NewCoins(sdk.NewCoin("uatom", math.NewInt(50000000000000))))
	suite.bankKeeper.SendCoinsFromModuleToAccount(
		suite.ctx,
		types.ModuleName,
		creator,
		sdk.NewCoins(sdk.NewCoin("uatom", math.NewInt(10000000000000))))

	// 1. create a test pool
	addr, _ := sdk.AccAddressFromBech32("meme10pzlhf8gz2c897djmdu5nsvjrym4tvy2qujsed")
	fmt.Printf("Initial minting: 1 ATOM\n")
	poolID, err := suite.keeper.CreateNewPool(
		suite.ctx,
		&types.MsgCreatePool{
			TokenDenom:  "udoge",
			Amount:      "100000",
			Creator:     "meme1j49359c08v66rthaqutfh6wc7k04sssp843kqj",
			NativeDenom: "uatom",
			Value: &sdk.Coin{
				Denom:  "uatom",
				Amount: math.NewInt(2000000),
			},
			IsNativeIn: true,
		},
		addr,
	)
	require.NoError(suite.T(), err)
	fmt.Printf("New created pool: %v\n", poolID)

	b := suite.bankKeeper.GetBalance(suite.ctx, creator, "udoge")
	fmt.Printf("Intial user DOGE balance: %v\n", b.String())

	pool, err := suite.keeper.GetPool(suite.ctx, poolID)
	require.NoError(suite.T(), err)

	fmt.Printf("Create pool ID: %v\n", pool.Id)

	userAddr, _ := sdk.AccAddressFromBech32("meme1cmpcwzqxp2354p6galq0cj37tk8rr7q3l8393n")
	fmt.Printf("Swap 0.1ATOM to DOGE\n")
	res, err := suite.keeper.SwapExactIn(
		suite.ctx,
		poolID,
		big.NewInt(100000),
		big.NewInt(100000),
		true,
		userAddr,
		sdk.NewCoin("uatom", math.NewInt(100000)),
	)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
	}
	require.NoError(suite.T(), err)
	fmt.Printf("You get %v (/10^7)\n", res)
}

func (suite *KeeperTestSuite) TestSwapExactOut() {

}
