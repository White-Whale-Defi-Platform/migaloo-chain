package v4_test

import (
	"cosmossdk.io/math"
	"fmt"
	"testing"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app/params"
	v4 "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app/upgrades/v4_1_0"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	apptesting "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app"
	"github.com/stretchr/testify/suite"
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

// Ensures the test does not error out.
func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup(s.T())
	// == CREATE MOCK VESTING ACCOUNT ==
	cVesting, unvested := v4.CreateMainnetVestingAccount(s.Ctx, s.App.BankKeeper, s.App.AccountKeeper)
	vestingAddr := cVesting.GetAddress()
	fmt.Printf("VestingAddr unvested: %+v\n", unvested)

	accVestingBalance := s.App.BankKeeper.GetAllBalances(s.Ctx, vestingAddr)
	fmt.Printf("Acc vesting bal: %s\n", accVestingBalance)

	// create many validators to confirm the unbonding code works
	newVal1 := s.SetupValidator(stakingtypes.Bonded)
	newVal2 := s.SetupValidator(stakingtypes.Bonded)
	newVal3 := s.SetupValidator(stakingtypes.Bonded)

	// Delegate tokens of the vesting multisig account
	s.StakingHelper.Delegate(vestingAddr, newVal1, sdk.NewInt(100))
	s.StakingHelper.Delegate(vestingAddr, newVal2, sdk.NewInt(200))
	s.StakingHelper.Delegate(vestingAddr, newVal3, sdk.NewInt(300))

	// Undelegate part of the tokens from val2 (test instant unbonding on undelegation started before upgrade)
	s.StakingHelper.Undelegate(vestingAddr, newVal3, sdk.NewInt(10), true)

	// Redelegate part of the tokens from val2 -> val3 (test instant unbonding on redelegations started before upgrade)
	_, err := s.App.StakingKeeper.BeginRedelegation(s.Ctx, vestingAddr, newVal2, newVal3, sdk.NewDec(1))
	s.Require().NoError(err)

	// Confirm delegated to 3 validators
	s.Require().Equal(3, len(s.App.StakingKeeper.GetAllDelegatorDelegations(s.Ctx, vestingAddr)))

	// == UPGRADE ==
	upgradeHeight := int64(5)
	s.ConfirmUpgradeSucceeded(v4.UpgradeName, upgradeHeight)

	// == VERIFICATION FEEBURN ==
	feeBurnParam := s.App.FeeBurnKeeper.GetParams(s.Ctx)
	s.Require().Equal("0", feeBurnParam.GetTxFeeBurnPercent())

	// VERIFY MULTISIGN MIGRATION
	accAfter := s.App.AccountKeeper.GetAccount(s.Ctx, vestingAddr)
	_, ok := accAfter.(*vestingtypes.ContinuousVestingAccount)
	s.Require().True(ok)

	s.Require().Equal(1, len(s.App.BankKeeper.GetAllBalances(s.Ctx, vestingAddr)))
	s.Require().Equal(4, len(s.App.StakingKeeper.GetAllDelegatorDelegations(s.Ctx, vestingAddr)))
	s.Require().Equal(0, len(s.App.StakingKeeper.GetRedelegations(s.Ctx, vestingAddr, 65535)))

	// check old multisign address balance
	oldMultisigBalance := s.App.BankKeeper.GetAllBalances(s.Ctx, sdk.MustAccAddressFromBech32(v4.NotionalMultisigVestingAccount))
	fmt.Printf("Old multisign address Upgrade Balance: %s\n", oldMultisigBalance)
	s.Require().True(oldMultisigBalance.AmountOf(params.BaseDenom).GTE(unvested))

	// check new multisign address balance
	newBalance := s.App.BankKeeper.GetAllBalances(s.Ctx, sdk.MustAccAddressFromBech32(v4.NewNotionalMultisigAccount))
	vestedBalance := cVesting.GetVestedCoins(s.Ctx.BlockTime())
	fmt.Printf("New multisign Upgrade Balance: %s, vestedBalance %s\n", newBalance, vestedBalance)
	s.Require().True(newBalance.AmountOf(params.BaseDenom).LTE(vestedBalance.AmountOf(params.BaseDenom)))
}

func (s *UpgradeTestSuite) TestMath() {
	s.Require().Equal(math.NewInt(7), math.NewInt(76).Quo(math.NewInt(10)))
	s.Require().Equal(math.NewInt(7), math.NewInt(79).Quo(math.NewInt(10)))
	s.Require().Equal(math.NewInt(1), math.NewInt(5).Quo(math.NewInt(3)))
}
