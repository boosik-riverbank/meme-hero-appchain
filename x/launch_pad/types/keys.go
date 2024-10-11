package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "launch_pad"
	StoreKey   = ModuleName
)

var (
	KeyPrefixPools = []byte{0x01}
)

func GetKeyPrefixPools(poolId uint64) []byte {
	return append(KeyPrefixPools, sdk.Uint64ToBigEndian(poolId)...)
}
