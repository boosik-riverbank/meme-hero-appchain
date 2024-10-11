package launch_pad

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/keeper"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/types"
)

var (
	_ module.AppModuleBasic = (*AppModule)(nil)
)

type AppModuleBasic struct{}

type AppModule struct {
	AppModuleBasic
	keeper     keeper.Keeper
	paramSpace paramtypes.Subspace
}

func NewAppModule(k keeper.Keeper, paramSpace paramtypes.Subspace) AppModule {
	return AppModule{
		keeper:     k,
		paramSpace: paramSpace,
	}
}

func (AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) RegisterServices(cfg module.Configurator) {

}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	am.keeper.InitGenesis(ctx, &genesisState)
}
