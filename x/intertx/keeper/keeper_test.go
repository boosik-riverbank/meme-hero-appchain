package keeper_test

import (
	"context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/interchain-security/v6/x/ccv/types"
	"github.com/cosmos/interchain-security/v6/x/intertx/keeper"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite
	cdc codec.Codec

	ctx context.Context

	keeper keeper.Keeper

	authKeeper   types.AccountKeeper
	scopedKeeper types.ScopedKeeper
	icaControllerKeeper
}
