package v4

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// CreateUpgradeHandler that migrates the chain from v4.2.5 to v4.2.6
func CreateUpgradeHandler(
	mm *module.Manager,
	bankKeeper bankKeeper.Keeper,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		// ignore the error if any
		if err := migrateFundFromDeadContacts(sdkCtx, bankKeeper); err != nil {
			sdkCtx.Logger().Error("migrateFundFromDeadContacts", "error", err)
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

// migrate fund from dead contacts to foundation
func migrateFundFromDeadContacts(
	ctx sdk.Context,
	bankKeeper bankKeeper.Keeper,
) error {
	deadContractAddr := sdk.MustAccAddressFromBech32(DeadContract)
	foundationAddr := sdk.MustAccAddressFromBech32(Foundation)

	// transfer token from dead contract to foundation

	// Get all balances from the dead contract
	allBalances := bankKeeper.GetAllBalances(ctx, deadContractAddr)
	if allBalances.IsZero() {
		return nil // Return early if no balances
	}

	return bankKeeper.SendCoins(ctx, deadContractAddr, foundationAddr, allBalances)
}
