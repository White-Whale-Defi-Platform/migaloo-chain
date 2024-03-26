package e2e

import (
	"testing"

	wasmibctesting "github.com/CosmWasm/wasmd/x/wasm/ibctesting"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app"
	tmtypes "github.com/cometbft/cometbft/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// DefaultMigalooAppFactory instantiates and sets up the default Migaloo app
func DefaultMigalooAppFactory(t *testing.T, valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, chainID string, opts []wasmkeeper.Option, balances ...banktypes.Balance) wasmibctesting.ChainApp {
	t.Helper()
	return app.SetupWithGenesisValSet(t, valSet, genAccs, chainID, opts, balances...)
}
