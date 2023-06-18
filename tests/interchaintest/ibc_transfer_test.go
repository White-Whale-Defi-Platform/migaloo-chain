package interchaintest

import (
	"context"
	"fmt"
	"testing"

	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestMigalooGaiaIBCTransfer spins up a Migaloo and Gaia network, initializes an IBC connection between them,
// and sends an ICS20 token transfer from Migaloo->Gaia and then back from Gaia->Migaloo.
func TestMigalooGaiaIBCTransfer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	ctx := context.Background()

	// Create chain factory with Migaloo and Gaia
	numVals := 1
	numFullNodes := 1

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          "Migaloo",
			ChainConfig:   migalooConfig,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          "gaia",
			Version:       "v9.0.2",
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
	})

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	migaloo, gaia := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)
	pathMigalooGaia := "migaloo-gaia"
	// Create relayer factory to utilize the go-relayer
	client, network := interchaintest.DockerSetup(t)

	r := interchaintest.NewBuiltinRelayerFactory(ibc.CosmosRly, zaptest.NewLogger(t)).Build(t, client, network)

	// Create a new Interchain object which describes the chains, relayers, and IBC connections we want to use
	ic := interchaintest.NewInterchain().
		AddChain(migaloo).
		AddChain(gaia).
		AddRelayer(r, "rly").
		AddLink(interchaintest.InterchainLink{
			Chain1:  migaloo,
			Chain2:  gaia,
			Relayer: r,
			Path:    pathMigalooGaia,
		})

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	err = ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: false,

		// This can be used to write to the block database which will index all block data e.g. txs, msgs, events, etc.
		// BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathMigalooGaia))
	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occurred while stopping the relayer: %s", err))
			}
		},
	)
	genesisWalletAmount := int64(10_000_000)
	// Create some user accounts on both chains
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, migaloo, gaia)

	// Wait a few blocks for relayer to start and for user accounts to be created
	err = testutil.WaitForBlocks(ctx, 5, migaloo, gaia)
	require.NoError(t, err)

	// Get our Bech32 encoded user addresses
	migalooUser, gaiaUser := users[0], users[1]

	migalooUserAddr := migalooUser.FormattedAddress()
	gaiaUserAddr := gaiaUser.FormattedAddress()

	// Get original account balances
	migalooOrigBal, err := migaloo.GetBalance(ctx, migalooUserAddr, migaloo.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, migalooOrigBal)

	gaiaOrigBal, err := gaia.GetBalance(ctx, gaiaUserAddr, gaia.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, gaiaOrigBal)

	// Compose an IBC transfer and send from migaloo -> Gaia
	const transferAmount = int64(1_000)
	transfer := ibc.WalletAmount{
		Address: gaiaUserAddr,
		Denom:   migaloo.Config().Denom,
		Amount:  transferAmount,
	}

	channel, err := ibc.GetTransferChannel(ctx, r, eRep, migaloo.Config().ChainID, gaia.Config().ChainID)
	require.NoError(t, err)

	transferTx, err := migaloo.SendIBCTransfer(ctx, channel.ChannelID, migalooUserAddr, transfer, ibc.TransferOptions{})
	require.NoError(t, err)

	migalooHeight, err := migaloo.Height(ctx)
	require.NoError(t, err)

	// Poll for the ack to know the transfer was successful
	_, err = testutil.PollForAck(ctx, migaloo, migalooHeight, migalooHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Get the IBC denom for stake on Gaia
	migalooTokenDenom := transfertypes.GetPrefixedDenom(channel.Counterparty.PortID, channel.Counterparty.ChannelID, migaloo.Config().Denom)
	migalooIBCDenom := transfertypes.ParseDenomTrace(migalooTokenDenom).IBCDenom()

	// Assert that the funds are no longer present in user acc on migaloo and are in the user acc on Gaia
	migalooUpdateBal, err := migaloo.GetBalance(ctx, migalooUserAddr, migaloo.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, migalooOrigBal-transferAmount, migalooUpdateBal)

	gaiaUpdateBal, err := gaia.GetBalance(ctx, gaiaUserAddr, migalooIBCDenom)
	require.NoError(t, err)
	require.Equal(t, transferAmount, gaiaUpdateBal)

	// Compose an IBC transfer and send from Gaia -> migaloo
	transfer = ibc.WalletAmount{
		Address: migalooUserAddr,
		Denom:   migalooIBCDenom,
		Amount:  transferAmount,
	}

	transferTx, err = gaia.SendIBCTransfer(ctx, channel.Counterparty.ChannelID, gaiaUserAddr, transfer, ibc.TransferOptions{})
	require.NoError(t, err)

	gaiaHeight, err := gaia.Height(ctx)
	require.NoError(t, err)

	// Poll for the ack to know the transfer was successful
	_, err = testutil.PollForAck(ctx, gaia, gaiaHeight, gaiaHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Assert that the funds are now back on migaloo and not on Gaia
	migalooUpdateBal, err = migaloo.GetBalance(ctx, migalooUserAddr, migaloo.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, migalooOrigBal, migalooUpdateBal)

	gaiaUpdateBal, err = gaia.GetBalance(ctx, gaiaUserAddr, migalooIBCDenom)
	require.NoError(t, err)
	require.Equal(t, int64(0), gaiaUpdateBal)
}
