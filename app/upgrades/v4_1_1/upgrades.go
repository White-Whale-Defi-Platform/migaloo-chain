package v4

import (
	feeburnkeeper "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/x/feeburn/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// CreateUpgradeHandler small security fix, can be a no-op, running mm.RunMigarions just to be sure
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	accountKeeper authkeeper.AccountKeeper,
	feeBurnKeeper feeburnkeeper.Keeper,

) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

		// Burning module permissions
		moduleAccI := accountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName)
		moduleAcc := moduleAccI.(*authtypes.ModuleAccount)
		moduleAcc.Permissions = []string{authtypes.Burner}
		accountKeeper.SetModuleAccount(ctx, moduleAcc)

		// set default fee_burn_percent to 50
		feeBurnParams := feeBurnKeeper.GetParams(ctx)
		feeBurnParams.TxFeeBurnPercent = "50"

		_ = feeBurnKeeper.SetParams(ctx, feeBurnParams)
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
