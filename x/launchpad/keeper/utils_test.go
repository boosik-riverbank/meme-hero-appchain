package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/launchpad/types"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestGenerateNewAddress(t *testing.T) {
	addr := GenerateNewAddress(types.ModuleName, "pool", sdk.Uint64ToBigEndian(1))
	fmt.Printf("addr: %v\n", addr)
}

func TestMulDiv(t *testing.T) {
	type testCase struct {
		a           *big.Int
		b           *big.Int
		denominator *big.Int
		expected    string
	}

	testData := []testCase{
		{
			a:           big.NewInt(1000),
			b:           big.NewInt(2000),
			denominator: big.NewInt(3),
			expected:    "666667",
		},
		{
			a:           big.NewInt(937813),
			b:           big.NewInt(152131),
			denominator: big.NewInt(7),
			expected:    "20381489929",
		},
		{
			a:           big.NewInt(12561028318),
			b:           big.NewInt(283719283721),
			denominator: big.NewInt(4),
			expected:    "890951489295539352820",
		},
	}

	for _, v := range testData {
		res, _ := MulDivRoundingUp(v.a, v.b, v.denominator)
		// answer, _ := new(big.Int).SetString(v.expected, 10)
		fmt.Printf("%v\n", res)

		require.Equal(t, res.String(), v.expected)
	}

}
