package keeper_test

import (
	"fmt"
	"math/big"
)

func (suite *KeeperTestSuite) TestQuoteAmountIn() {
	// 2	1005937679
	out, err := suite.keeper.QuoteAmountIn(big.NewInt(1_000_0000_000000), big.NewInt(1005937679_000000), big.NewInt(32_000_000))
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
	}
	fmt.Printf("%v\n", out)
}

func (suite *KeeperTestSuite) TestQuoteAmountOut() {
	out, err := suite.keeper.QuoteAmountOut(big.NewInt(1_000_000), big.NewInt(1073000191_000_000), big.NewInt(30_000_000))
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
	}
	fmt.Printf("%v\n", out)
}

func (suite *KeeperTestSuite) TestBuyToken() {

}
