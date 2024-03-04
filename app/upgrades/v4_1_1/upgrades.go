package v4

import (
	feeburnkeeper "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/x/feeburn/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func UpdateAccountPermissionAndFeeBurnPercent(ctx sdk.Context, fk feeburnkeeper.Keeper, ak authkeeper.AccountKeeper) {
	// Burning module permissions
	moduleAccI := ak.GetModuleAccount(ctx, authtypes.FeeCollectorName)
	moduleAcc := moduleAccI.(*authtypes.ModuleAccount)
	moduleAcc.Permissions = []string{authtypes.Burner}
	ak.SetModuleAccount(ctx, moduleAcc)

	// set default fee_burn_percent to 10
	feeBurnParams := fk.GetParams(ctx)
	feeBurnParams.TxFeeBurnPercent = "10"

	_ = fk.SetParams(ctx, feeBurnParams)
}
