package v2_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/app/apptesting"
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
