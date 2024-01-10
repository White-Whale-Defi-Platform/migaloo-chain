package v4_test

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"testing"

	apptesting "github.com/White-Whale-Defi-Platform/migaloo-chain/v3/app"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
)

const (
	v4UpgradeHeight = int64(10)
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (suite *UpgradeTestSuite) TestUpgrade() {
	suite.Setup(suite.T(), apptesting.SimAppChainID)
	dummyUpgrade(suite)
	feeBurnParam := suite.App.FeeBurnKeeper.GetParams(suite.Ctx)
	suite.Require().Equal("0", feeBurnParam.GetTxFeeBurnPercent())
}

func dummyUpgrade(s *UpgradeTestSuite) {
	s.Ctx = s.Ctx.WithBlockHeight(v4UpgradeHeight - 1)
	plan := upgradetypes.Plan{Name: "v4.1.0", Height: v4UpgradeHeight}
	err := s.App.UpgradeKeeper.ScheduleUpgrade(s.Ctx, plan)
	s.Require().NoError(err)
	_, exists := s.App.UpgradeKeeper.GetUpgradePlan(s.Ctx)
	s.Require().True(exists)

	s.Ctx = s.Ctx.WithBlockHeight(v4UpgradeHeight)

	s.Require().NotPanics(func() {
		beginBlockRequest := abci.RequestBeginBlock{}
		s.App.BeginBlocker(s.Ctx, beginBlockRequest)
	})
}
