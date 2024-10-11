package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec

	paramSpace paramtypes.Subspace

	feeCollector sdk.AccAddress

	// Data
	length uint64
	pools  collections.Map[uint64, types.BondingCurvePool]

	// Keepers
	authKeeper types.AuthKeeper
	bankKeeper types.BankKeeper
}

func NewKeeper(
	storeService store.KVStoreService,
	cdc codec.BinaryCodec,
	feeCollector sdk.AccAddress,
	authKeeper types.AuthKeeper,
	bankKeeper types.BankKeeper) Keeper {
	sb := collections.NewSchemaBuilder(storeService)
	return Keeper{
		length:       0,
		feeCollector: feeCollector,
		authKeeper:   authKeeper,
		bankKeeper:   bankKeeper,
		pools:        collections.NewMap(sb, types.KeyPrefixPools, "pools", collections.Uint64Key, codec.CollValue[types.BondingCurvePool](cdc)),
	}
}
