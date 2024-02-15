package v4_test

import (
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
	newVal4 := s.SetupValidator(stakingtypes.Bonded)
	newVal5 := s.SetupValidator(stakingtypes.Bonded)
	newVal6 := s.SetupValidator(stakingtypes.Bonded)
	newVal7 := s.SetupValidator(stakingtypes.Bonded)
	newVal8 := s.SetupValidator(stakingtypes.Bonded)
	newVal9 := s.SetupValidator(stakingtypes.Bonded)
	newVal10 := s.SetupValidator(stakingtypes.Bonded)
	newVal11 := s.SetupValidator(stakingtypes.Bonded)
	newVal12 := s.SetupValidator(stakingtypes.Bonded)

	// Delegate tokens of the vesting multisig account
	s.StakingHelper.Delegate(vestingAddr, newVal1, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal2, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal3, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal4, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal5, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal6, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal7, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal8, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal9, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal10, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal11, sdk.NewInt(300))
	s.StakingHelper.Delegate(vestingAddr, newVal12, sdk.NewInt(300))

	// Undelegate part of the tokens from val2 (test instant unbonding on undelegation started before upgrade)
	s.StakingHelper.Undelegate(vestingAddr, newVal3, sdk.NewInt(10), true)

	// Redelegate part of the tokens from val2 -> val3 (test instant unbonding on redelegations started before upgrade)
	_, err := s.App.StakingKeeper.BeginRedelegation(s.Ctx, vestingAddr, newVal2, newVal3, sdk.NewDec(1))
	s.Require().NoError(err)

	// Confirm delegated to 12 validators
	s.Require().Equal(12, len(s.App.StakingKeeper.GetAllDelegatorDelegations(s.Ctx, vestingAddr)))

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

	s.Require().Equal(0, len(s.App.BankKeeper.GetAllBalances(s.Ctx, vestingAddr)))
	// now delegated to top 10 validator
	s.Require().Equal(10, len(s.App.StakingKeeper.GetAllDelegatorDelegations(s.Ctx, vestingAddr)))
	s.Require().Equal(0, len(s.App.StakingKeeper.GetRedelegations(s.Ctx, vestingAddr, 65535)))

	// check old multisign address balance
	oldMultisigBalance := s.App.BankKeeper.GetAllBalances(s.Ctx, sdk.MustAccAddressFromBech32(v4.NotionalMultisigVestingAccount))
	fmt.Printf("Old multisign address Upgrade Balance: %s\n", oldMultisigBalance)
	s.Require().True(oldMultisigBalance.Empty())
	totalDelegateBalance := s.App.StakingKeeper.GetDelegatorBonded(s.Ctx, sdk.MustAccAddressFromBech32(v4.NotionalMultisigVestingAccount))
	fmt.Printf("old multisign address totalDelegateBalance %v\n", totalDelegateBalance)
	s.Require().True(totalDelegateBalance.Equal(unvested))

	// check new multisign address balance
	newBalance := s.App.BankKeeper.GetAllBalances(s.Ctx, sdk.MustAccAddressFromBech32(v4.NewNotionalMultisigAccount))
	vestedBalance := cVesting.GetVestedCoins(s.Ctx.BlockTime())
	fmt.Printf("New multisign Upgrade Balance: %s, vestedBalance %s\n", newBalance, vestedBalance)
	s.Require().True(vestedBalance.AmountOf(params.BaseDenom).Equal(newBalance.AmountOf(params.BaseDenom)))
}
