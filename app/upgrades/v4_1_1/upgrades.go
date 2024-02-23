package v4

import (
	"fmt"
	"time"

	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	clientkeeper "github.com/cosmos/ibc-go/v7/modules/core/02-client/keeper"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
)

// CreateUpgradeHandler small security fix, can be a no-op, running mm.RunMigarions just to be sure
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	clientKeeper clientkeeper.Keeper,
	paramsKeeper paramskeeper.Keeper,
	consensusParamsKeeper consensuskeeper.Keeper,
	icacontrollerKeeper icacontrollerkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	stakingKeeper stakingKeeper.Keeper,
	bankKeeper bankKeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// READ: https://github.com/cosmos/cosmos-sdk/blob/v0.47.4/UPGRADING.md#xconsensus
		baseAppLegacySS := paramsKeeper.Subspace(baseapp.Paramspace).
			WithKeyTable(paramstypes.ConsensusParamsKeyTable())
		baseapp.MigrateParams(ctx, baseAppLegacySS, &consensusParamsKeeper)

		// READ: https://github.com/cosmos/ibc-go/blob/v7.2.0/docs/migrations/v7-to-v7_1.md#chains
		params := clientKeeper.GetParams(ctx)
		params.AllowedClients = append(params.AllowedClients, ibcexported.Localhost)
		clientKeeper.SetParams(ctx, params)

		// READ: https://github.com/terra-money/core/issues/166
		icacontrollerKeeper.SetParams(ctx, icacontrollertypes.DefaultParams())

		// Burning module permissions
		moduleAccI := accountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName)
		moduleAcc := moduleAccI.(*authtypes.ModuleAccount)
		moduleAcc.Permissions = []string{authtypes.Burner}
		accountKeeper.SetModuleAccount(ctx, moduleAcc)

		migrateMultisigVesting(ctx, stakingKeeper, bankKeeper, accountKeeper)
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

// migrateMultisigVesting moves the unvested from the ContinuousVestingAccount -> the new ContinuousVestingAccount
// - Retrieves the old multisig vesting account
// - Instantly finish all redelegations, then unbond all tokens.
// - Transfer unvested token to the new multisign ContinuousVestingAccount
func migrateMultisigVesting(ctx sdk.Context,
	stakingKeeper stakingKeeper.Keeper,
	bankKeeper bankKeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
) {
	currentAddr := sdk.MustAccAddressFromBech32(NotionalMultisigVestingAccount)
	newAddr := sdk.MustAccAddressFromBech32(NewNotionalMultisigAccount)

	currentAcc := accountKeeper.GetAccount(ctx, currentAddr)

	currentVestingAcc, ok := currentAcc.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		// skip if account invalid
		fmt.Printf("err currentAcc.(*vestingtypes.ContinuousVestingAccount): %+v\n", currentAcc)
		return
	}
	// process migrate
	processMigrateMultisig(ctx, stakingKeeper, bankKeeper, accountKeeper, currentAddr, newAddr, currentVestingAcc)
}

func processMigrateMultisig(ctx sdk.Context, stakingKeeper stakingKeeper.Keeper,
	bankKeeper bankKeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	currentAddr, newAddr sdk.AccAddress,
	oldAcc *vestingtypes.ContinuousVestingAccount,
) {
	redelegated, err := completeAllRedelegations(ctx, ctx.BlockTime(), stakingKeeper, currentAddr)
	if err != nil {
		panic(err)
	}

	unbonded, err := unbondAllAndFinish(ctx, ctx.BlockTime(), stakingKeeper, currentAddr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("currentAddr Instant Redelegations: %s\n", redelegated)
	fmt.Printf("currentAddr Instant Unbonding: %s\n", unbonded)

	// get vesting balance
	vestingCoin := oldAcc.GetVestingCoins(ctx.BlockTime())
	fmt.Printf("Total vesting balance %v\n", vestingCoin)
	if vestingCoin.Empty() {
		panic(err)
	}

	// setup NewContinuousVestingAccount
	newAccount := authtypes.NewBaseAccountWithAddress(newAddr)
	newVestingAcc := vestingtypes.NewContinuousVestingAccount(newAccount, vestingCoin,
		ctx.BlockTime().Unix(), oldAcc.EndTime)
	accountKeeper.SetAccount(ctx, newVestingAcc)

	// end current ContinuousVestingAccount
	oldAcc.OriginalVesting = oldAcc.OriginalVesting.Sub(vestingCoin[0])
	oldAcc.DelegatedFree = sdk.NewCoins()
	oldAcc.DelegatedVesting = sdk.NewCoins()
	oldAcc.EndTime = ctx.BlockTime().Unix()
	accountKeeper.SetAccount(ctx, oldAcc)

	// move vesting coin to new ContinuousVestingAccount
	err = bankKeeper.SendCoins(ctx, currentAddr, newAddr, vestingCoin)
	if err != nil {
		panic(err)
	}
}

func GetVestingCoin(ctx sdk.Context, acc *vestingtypes.ContinuousVestingAccount) (unvested math.Int) {
	vestingCoin := acc.GetVestingCoins(ctx.BlockTime())
	return vestingCoin[0].Amount
}

func completeAllRedelegations(ctx sdk.Context, now time.Time,
	stakingKeeper stakingKeeper.Keeper,
	accAddr sdk.AccAddress,
) (math.Int, error) {
	redelegatedAmt := math.ZeroInt()

	for _, activeRedelegation := range stakingKeeper.GetRedelegations(ctx, accAddr, 65535) {
		redelegationSrc, _ := sdk.ValAddressFromBech32(activeRedelegation.ValidatorSrcAddress)
		redelegationDst, _ := sdk.ValAddressFromBech32(activeRedelegation.ValidatorDstAddress)

		// set all entry completionTime to now so we can complete re-delegation
		for i := range activeRedelegation.Entries {
			activeRedelegation.Entries[i].CompletionTime = now
			redelegatedAmt = redelegatedAmt.Add(math.Int(activeRedelegation.Entries[i].SharesDst))
		}

		stakingKeeper.SetRedelegation(ctx, activeRedelegation)
		_, err := stakingKeeper.CompleteRedelegation(ctx, accAddr, redelegationSrc, redelegationDst)
		if err != nil {
			return redelegatedAmt, err
		}
	}

	return redelegatedAmt, nil
}

func unbondAllAndFinish(ctx sdk.Context, now time.Time,
	stakingKeeper stakingKeeper.Keeper,
	accAddr sdk.AccAddress,
) (math.Int, error) {
	unbondedAmt := math.ZeroInt()

	// Unbond all delegations from the account
	for _, delegation := range stakingKeeper.GetAllDelegatorDelegations(ctx, accAddr) {
		validatorValAddr := delegation.GetValidatorAddr()
		_, found := stakingKeeper.GetValidator(ctx, validatorValAddr)
		if !found {
			continue
		}

		_, err := stakingKeeper.Undelegate(ctx, accAddr, validatorValAddr, delegation.GetShares())
		if err != nil {
			return math.ZeroInt(), err
		}
	}

	// Take all unbonding and complete them.
	for _, unbondingDelegation := range stakingKeeper.GetAllUnbondingDelegations(ctx, accAddr) {
		validatorStringAddr := unbondingDelegation.ValidatorAddress
		validatorValAddr, _ := sdk.ValAddressFromBech32(validatorStringAddr)

		// Complete unbonding delegation
		for i := range unbondingDelegation.Entries {
			unbondingDelegation.Entries[i].CompletionTime = now
			unbondedAmt = unbondedAmt.Add(unbondingDelegation.Entries[i].Balance)
		}

		stakingKeeper.SetUnbondingDelegation(ctx, unbondingDelegation)
		_, err := stakingKeeper.CompleteUnbonding(ctx, accAddr, validatorValAddr)
		if err != nil {
			return math.ZeroInt(), err
		}
	}

	return unbondedAmt, nil
}
