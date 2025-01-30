package v4

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

// CreateUpgradeHandler that migrates the chain from v4.2.4 to v4.2.5
func CreateUpgradeHandler(
	mm *module.Manager,
	_ *stakingKeeper.Keeper,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
