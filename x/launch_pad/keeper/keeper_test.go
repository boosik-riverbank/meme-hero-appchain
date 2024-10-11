package keeper_test

import (
	"context"
	"cosmossdk.io/core/header"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/keeper"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/test_utils"
	"github.com/cosmos/interchain-security/v6/x/launch_pad/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var (
	launchPadAcc = authtypes.NewEmptyModuleAccount(types.ModuleName, authtypes.Minter)
)

type KeeperTestSuite struct {
	suite.Suite
	cdc codec.Codec

	ctx context.Context

	keeper     keeper.Keeper
	bankKeeper types.BankKeeper
	authKeeper *types.MockAuthKeeper

	queryClient types.QueryClient
	msgServer   types.MsgServer
	baseApp     *baseapp.BaseApp
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithHeaderInfo(header.Info{Time: time.Now()})

	fmt.Println("Set test context")
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("meme", "memepub")
	config.Seal()
	suite.ctx = ctx

	encCfg := moduletestutil.MakeTestEncodingConfig()
	suite.cdc = encCfg.Codec

	suite.baseApp = baseapp.NewBaseApp(
		"launch_pad",
		log.NewNopLogger(),
		testCtx.DB,
		encCfg.TxConfig.TxDecoder(),
	)
	suite.baseApp.SetCMS(testCtx.CMS)
	suite.baseApp.SetInterfaceRegistry(encCfg.InterfaceRegistry)

	ctrl := gomock.NewController(suite.T())
	bankKeeper := test_utils.NewMockBankKeeper()
	authKeeper := types.NewMockAuthKeeper(ctrl)
	suite.authKeeper = authKeeper
	suite.bankKeeper = bankKeeper

	storeService := runtime.NewKVStoreService(key)
	feeCollector, _ := sdk.AccAddressFromBech32("meme1cmpcwzqxp2354p6galq0cj37tk8rr7q3l8393n")

	suite.keeper = keeper.NewKeeper(
		storeService,
		suite.cdc,
		feeCollector,
		authKeeper,
		bankKeeper)
}

// BANK KEEPER
func (suite *KeeperTestSuite) mockSendCoinsFromModuleToAccount(moduleAcc *authtypes.ModuleAccount, _ sdk.AccAddress) {
	suite.authKeeper.EXPECT().GetModuleAddress(moduleAcc.Name).Return(moduleAcc.GetAddress())
	suite.authKeeper.EXPECT().GetAccount(suite.ctx, moduleAcc.GetAddress()).Return(moduleAcc)
}

//func (suite *KeeperTestSuite) mockGetBalance(address sdk.AccAddress, denom string, testReturn sdk.Coin) {
//	suite.bankKeeper.EXPECT().GetBalance(suite.ctx, address, denom).Return(testReturn)
//}
//
//func (suite *KeeperTestSuite) mockMintCoins(moduleName string, amt sdk.Coins) {
//	address := suite.authKeeper.GetModuleAddress(moduleName)
//	for _, coin := range amt {
//		// exist := suite.bankKeeper.GetBalance(suite.ctx, address, coin.Denom)
//		suite.bankKeeper.EXPECT().GetBalance(suite.ctx, address, coin.Denom).Return(coin)
//	}
//	suite.bankKeeper.EXPECT().MintCoins(suite.ctx, gomock.Any(), amt).Return(nil)
//}

// AUTH KEEPER
func (suite *KeeperTestSuite) mockGetModuleAddress(moduleName string, address sdk.AccAddress) {
	suite.authKeeper.EXPECT().GetModuleAddress(moduleName).Return(address)
}
