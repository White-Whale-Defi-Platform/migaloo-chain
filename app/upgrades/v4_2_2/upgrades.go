package v4

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

// CreateUpgradeHandler that migrates the chain from v4.2.0 to v4.2.1
func CreateUpgradeHandler(
	mm *module.Manager,
	sk *stakingKeeper.Keeper,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		stakingParams, err := sk.GetParams(sdkCtx)
		if err != nil {
			panic(err)
		}

		stakingParams.MaxValidators = 45
		err = sk.SetParams(ctx, stakingParams)
		if err != nil {
			panic(err)
		}
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
