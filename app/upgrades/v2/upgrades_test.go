package v2_test

import (
	"testing"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/app/apptesting"
	"github.com/stretchr/testify/suite"
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func (suite *UpgradeTestSuite) SetupTest() {
	suite.Setup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

const dummyUpgradeHeight = 5

func (suite *UpgradeTestSuite) TestUpgrade() {
	suite.Setup()
	suite.ConfirmUpgradeSucceededs("v2", dummyUpgradeHeight)
}
