package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/intertx/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	// TODO : implement this!
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	// TODO : implement this!
	return nil
}
