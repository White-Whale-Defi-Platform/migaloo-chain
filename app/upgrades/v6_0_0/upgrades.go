package v6

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// CreateUpgradeHandler creates the upgrade handler for v6.0.0
// This upgrade consolidates all chain assets to a custody wallet for the sunset process.
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	accountKeeper accountKeeper.AccountKeeper,
	bankKeeper bankKeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("Starting v6.0.0 upgrade: Asset consolidation for sunset")

		// Consolidate all assets to custody wallet
		if err := consolidateAssets(ctx, accountKeeper, bankKeeper); err != nil {
			// Log error but don't fail the upgrade - some accounts may have issues
			ctx.Logger().Error("Error during asset consolidation", "error", err)
		}

		ctx.Logger().Info("Completed v6.0.0 upgrade")
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

// consolidateAssets moves all balances from all accounts to the custody wallet
func consolidateAssets(
	ctx sdk.Context,
	accountKeeper accountKeeper.AccountKeeper,
	bankKeeper bankKeeper.Keeper,
) error {
	custodyAddr := sdk.MustAccAddressFromBech32(CustodyWallet)

	var transferCount int
	var contractCount int
	var totalCoins sdk.Coins
	var errors []error

	// Iterate through all accounts (includes regular accounts, contracts, etc.)
	accountKeeper.IterateAccounts(ctx, func(account authtypes.AccountI) bool {
		addr := account.GetAddress()
		accountType := reflect.TypeOf(account).String()

		// Skip the custody wallet itself
		if addr.Equals(custodyAddr) {
			return false // continue iteration
		}

		// Only skip critical module accounts (ones that would break chain consensus)
		// Transfer from other module accounts (alliance, distribution, rewards, etc.)
		if moduleAcc, ok := account.(*authtypes.ModuleAccount); ok {
			criticalModules := map[string]bool{
				"bonded_tokens_pool":     true, // Staking pool - critical
				"not_bonded_tokens_pool": true, // Staking pool - critical
				"gov":                    true, // Governance - critical
				"mint":                   true, // Token minting - critical
			}

			if criticalModules[moduleAcc.Name] {
				ctx.Logger().Debug("Skipping critical module account", "address", addr.String(), "name", moduleAcc.Name)
				return false // continue iteration
			}

			// Non-critical module account - transfer funds
			ctx.Logger().Info("Transferring from non-critical module account", "address", addr.String(), "name", moduleAcc.Name)
		}

		// Get all balances for this account
		balances := bankKeeper.GetAllBalances(ctx, addr)
		if balances.IsZero() {
			return false // continue iteration
		}

		// Transfer all balances to custody wallet
		if err := bankKeeper.SendCoins(ctx, addr, custodyAddr, balances); err != nil {
			ctx.Logger().Error("Failed to transfer from account",
				"from", addr.String(),
				"type", accountType,
				"amount", balances.String(),
				"error", err,
			)
			errors = append(errors, fmt.Errorf("failed to transfer from %s (%s): %w", addr.String(), accountType, err))
			return false // continue iteration despite error
		}

		transferCount++
		totalCoins = totalCoins.Add(balances...)

		// Check if this looks like a contract address (longer addresses, specific patterns)
		// CosmWasm contracts are typically BaseAccount but with contract-derived addresses
		addrStr := addr.String()
		if len(addrStr) > 50 {
			contractCount++
			ctx.Logger().Info("Transferred assets from CONTRACT to custody",
				"from", addrStr,
				"type", accountType,
				"amount", balances.String(),
			)
		} else {
			ctx.Logger().Info("Transferred assets to custody",
				"from", addrStr,
				"type", accountType,
				"amount", balances.String(),
			)
		}

		return false // continue iteration
	})

	ctx.Logger().Info("Asset consolidation complete",
		"transfers", transferCount,
		"contracts", contractCount,
		"total_coins", totalCoins.String(),
		"errors", len(errors),
	)

	if len(errors) > 0 {
		return fmt.Errorf("encountered %d errors during consolidation", len(errors))
	}

	return nil
}
