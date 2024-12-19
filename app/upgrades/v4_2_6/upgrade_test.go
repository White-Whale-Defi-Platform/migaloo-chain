package v4_test

import (
	"testing"

	"cosmossdk.io/math"
	v4 "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app/upgrades/v4_2_6"
	sdk "github.com/cosmos/cosmos-sdk/types"

	apptesting "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app"
	"github.com/stretchr/testify/suite"
)

const (
	MockDeadContractBalance = 1000000000000000000
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) MockBankBalances() {
	deadContractAddr := sdk.MustAccAddressFromBech32(v4.DeadContract)

	coins := sdk.NewCoins(
		sdk.NewCoin(v4.Denom, math.NewInt(MockDeadContractBalance)),
	)

	// Mint coins to the dead contract
	err := s.App.BankKeeper.MintCoins(s.Ctx, "mint", coins)
	s.Require().NoError(err)
	err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, "mint", deadContractAddr, coins)
	s.Require().NoError(err)

	// require the balance is correct
	balance := s.App.BankKeeper.GetAllBalances(s.Ctx, deadContractAddr)
	s.Require().Equal(coins, balance)
}

// Ensures the test does not error out.
func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup(s.T())
	// == CREATE MOCK VESTING ACCOUNT ==
	s.MockBankBalances()

	// == UPGRADE ==
	upgradeHeight := int64(5)

	// Execute upgrade
	s.ConfirmUpgradeSucceeded(v4.UpgradeName, upgradeHeight)

	// Dead contract balance get drained
	deadContractAddr := sdk.MustAccAddressFromBech32(v4.DeadContract)
	contractBalance := s.App.BankKeeper.GetAllBalances(s.Ctx, deadContractAddr)
	s.Require().Equal(int64(0), contractBalance.AmountOf(v4.Denom).Int64())

	// Foundation balance is increased
	foundationAddr := sdk.MustAccAddressFromBech32(v4.Foundation)
	foundationBalance := s.App.BankKeeper.GetAllBalances(s.Ctx, foundationAddr)
	s.T().Logf("balance: %v", foundationBalance)
	s.Require().Equal(int64(MockDeadContractBalance), foundationBalance.AmountOf(v4.Denom).Int64()) // Add int64 cast
}
