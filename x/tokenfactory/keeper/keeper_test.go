package keeper_test

//NOTE: this whole file needs to be refactored.

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/app"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/x/tokenfactory/keeper"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/x/tokenfactory/testhelpers"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/x/tokenfactory/types"
)

//TODO: refactor tests to match sdk 46 upstream

type KeeperTestSuite struct {
	app.KeeperTestHelper

	queryClient types.QueryClient
	msgServer   types.MsgServer
	// defaultDenom is on the suite, as it depends on the creator test address.
	defaultDenom string
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	// Fund every TestAcc with two denoms, one of which is the denom creation fee
	fundAccsAmount := sdk.NewCoins(sdk.NewCoin(types.DefaultParams().DenomCreationFee[0].Denom, types.DefaultParams().DenomCreationFee[0].Amount.MulRaw(100)), sdk.NewCoin(testhelpers.SecondaryDenom, testhelpers.SecondaryAmount))
	for _, acc := range suite.TestAccs {
		suite.FundAcc(acc, fundAccsAmount)
	}

	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.App.TokenFactoryKeeper)
}

func (suite *KeeperTestSuite) CreateDefaultDenom() {
	res, _ := suite.msgServer.CreateDenom(sdk.WrapSDKContext(suite.Ctx), types.NewMsgCreateDenom(suite.TestAccs[0].String(), "bitcoin"))
	suite.defaultDenom = res.GetNewTokenDenom()
}

func (suite *KeeperTestSuite) TestCreateModuleAccount() {
	app := suite.App

	// remove module account
	tokenfactoryModuleAccount := app.AccountKeeper.GetAccount(suite.Ctx, app.AccountKeeper.GetModuleAddress(types.ModuleName))
	app.AccountKeeper.RemoveAccount(suite.Ctx, tokenfactoryModuleAccount)

	// ensure module account was removed
	suite.Ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	tokenfactoryModuleAccount = app.AccountKeeper.GetAccount(suite.Ctx, app.AccountKeeper.GetModuleAddress(types.ModuleName))
	suite.Require().Nil(tokenfactoryModuleAccount)

	// create module account
	app.TokenFactoryKeeper.CreateModuleAccount(suite.Ctx)

	// check that the module account is now initialized
	tokenfactoryModuleAccount = app.AccountKeeper.GetAccount(suite.Ctx, app.AccountKeeper.GetModuleAddress(types.ModuleName))
	suite.Require().NotNil(tokenfactoryModuleAccount)
}
