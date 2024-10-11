package keeper_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/big"
)

func (suite *KeeperTestSuite) TestCalculateTokenOffset() {
	res, err := suite.keeper.CalculateTokenOffset(
		big.NewInt(460_000_000),
		big.NewInt(30_000_000),
		big.NewInt(8000),
		big.NewInt(1_000_000_000_000_000),
	)
	require.NoError(suite.T(), err)
	fmt.Printf("%v", res.String())
}
