package feeburn

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/keeper"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	err := k.SetParams(ctx, genState.Params)
	if err != nil {
		panic(errorsmod.Wrapf(err, "error setting params"))
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	return genesis
}
