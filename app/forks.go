package app

import (
	v4 "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app/upgrades/v4_1_2"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlockForks executes any necessary fork logic based upon the current block height.
func BeginBlockForks(ctx sdk.Context, app *MigalooApp) {
	if ctx.BlockHeight() == v4.UpgradeHeight {
		v4.UpdateAlliance(ctx, app.AllianceKeeper)
	}
}
