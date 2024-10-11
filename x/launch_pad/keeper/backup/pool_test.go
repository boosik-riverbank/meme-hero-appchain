package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/types"
)

func (suite *KeeperTestSuite) TestCreateNewPool() {
	require := suite.Require()
	addr, _ := sdk.AccAddressFromBech32("meme10pzlhf8gz2c897djmdu5nsvjrym4tvy2qujsed")
	creatorAddr, _ := sdk.AccAddressFromBech32("meme1j49359c08v66rthaqutfh6wc7k04sssp843kqj")
	id, err := suite.keeper.CreateNewPool(suite.ctx, &types.MsgCreatePool{
		TokenDenom:  "utest",
		Amount:      "10000",
		Creator:     creatorAddr.String(),
		NativeDenom: "",
	}, addr)
	require.NoError(err)
	require.Equal(id, uint64(1))

	result, err := suite.keeper.GetPool(suite.ctx, uint64(1))
	require.NoError(err)
	fmt.Printf("%v", result)
}

func (suite *KeeperTestSuite) TestDeletePool() {}

func (suite *KeeperTestSuite) TestGetPools() {}

func (suite *KeeperTestSuite) TestGetPool() {}

func (suite *KeeperTestSuite) TestGetPoolAmount() {}
