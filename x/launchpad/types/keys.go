package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "launchpad"
	StoreKey   = ModuleName
)

var (
	KeyPrefixPools     = []byte{0x01}
	KeyPrefixPoolCount = []byte{0x02}
)

func GetKeyPrefixPools(poolId uint64) []byte {
	return append(KeyPrefixPools, sdk.Uint64ToBigEndian(poolId)...)
}
