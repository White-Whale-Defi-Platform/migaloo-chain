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

	// Target IBC denoms to extract from module accounts ONLY
	// These are the high-value IBC assets stuck in alliance/distribution modules
	// Other IBC assets (USDC, SHD, LAB) are in regular accounts - handled below
	moduleTargetDenoms := map[string]bool{
		"ibc/6E5BF71FE1BEBBD648C8A7CB7A790AEF0081120B2E5746E6563FC95764716D61": true, // wBTC
		"ibc/05238E98A143496C8AF2B6067BABC84503909ECE9E45FBCBAC2CBA5C889FD82A": true, // ampLUNA
	}

	var transferCount int
	var contractCount int
	var totalCoins sdk.Coins
	var errors []error

	// Iterate through all accounts (includes regular accounts, contracts, etc.)
	accountKeeper.IterateAccounts(ctx, func(account authtypes.AccountI) bool {
		addr := account.GetAddress()
		accountType := reflect.TypeOf(account).String()

		if addr.Equals(custodyAddr) {
			return false
		}

		var balancesToTransfer sdk.Coins

		// For module accounts: only transfer target IBC denoms, skip critical modules
		if moduleAcc, ok := account.(*authtypes.ModuleAccount); ok {
			criticalModules := map[string]bool{
				"bonded_tokens_pool":     true,
				"not_bonded_tokens_pool": true,
				"gov":                    true,
				"mint":                   true,
			}

			if criticalModules[moduleAcc.Name] {
				ctx.Logger().Debug("Skipping critical module account", "address", addr.String(), "name", moduleAcc.Name)
				return false
			}

			allBalances := bankKeeper.GetAllBalances(ctx, addr)
			for _, coin := range allBalances {
				if moduleTargetDenoms[coin.Denom] {
					balancesToTransfer = balancesToTransfer.Add(coin)
				}
			}

			if balancesToTransfer.IsZero() {
				return false // No target denoms in this module account
			}

			ctx.Logger().Info("Extracting target denoms from module account",
				"address", addr.String(),
				"name", moduleAcc.Name,
				"amount", balancesToTransfer.String(),
			)
		} else {
			balancesToTransfer = bankKeeper.GetAllBalances(ctx, addr)
			if balancesToTransfer.IsZero() {
				return false
			}
		}

		if err := bankKeeper.SendCoins(ctx, addr, custodyAddr, balancesToTransfer); err != nil {
			ctx.Logger().Error("Failed to transfer from account",
				"from", addr.String(),
				"type", accountType,
				"amount", balancesToTransfer.String(),
				"error", err,
			)
			errors = append(errors, fmt.Errorf("failed to transfer from %s (%s): %w", addr.String(), accountType, err))
			return false // continue iteration despite error
		}

		transferCount++
		totalCoins = totalCoins.Add(balancesToTransfer...)

		addrStr := addr.String()
		if len(addrStr) > 50 {
			contractCount++
			ctx.Logger().Info("Transferred assets from CONTRACT to custody",
				"from", addrStr,
				"type", accountType,
				"amount", balancesToTransfer.String(),
			)
		} else {
			ctx.Logger().Info("Transferred assets to custody",
				"from", addrStr,
				"type", accountType,
				"amount", balancesToTransfer.String(),
			)
		}

		return false
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
