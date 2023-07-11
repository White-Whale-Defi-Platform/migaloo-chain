package interchaintest

import (
	"context"
	"fmt"
	"strings"
	"testing"

	helpers "github.com/White-Whale-Defi-Platform/migaloo-chain/tests/interchaintest/helpers"
	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	interchaintestrelayer "github.com/strangelove-ventures/interchaintest/v7/relayer"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestIBCHooks ensures the ibc-hooks middleware from osmosis works.
func TestIBCHooks(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Create chain factory with migaloo and migaloo2
	numVals := 1
	numFullNodes := 0

	cfg2 := migalooConfig.Clone()
	cfg2.Name = "migaloo-counterparty"
	cfg2.ChainID = "migaloo-counterparty-2"

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          "migaloo",
			ChainConfig:   migalooConfig,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          "migaloo",
			ChainConfig:   cfg2,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
	})

	const (
		path = "ibc-path"
	)

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	client, network := interchaintest.DockerSetup(t)

	migaloo, migaloo2 := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	relayerType, relayerName := ibc.CosmosRly, "relay"

	// Get a relayer instance
	rf := interchaintest.NewBuiltinRelayerFactory(
		relayerType,
		zaptest.NewLogger(t),
		interchaintestrelayer.CustomDockerImage(IBCRelayerImage, IBCRelayerVersion, "100:1000"),
		interchaintestrelayer.StartupFlags("--processor", "events", "--block-history", "100"),
	)

	r := rf.Build(t, client, network)

	ic := interchaintest.NewInterchain().
		AddChain(migaloo).
		AddChain(migaloo2).
		AddRelayer(r, relayerName).
		AddLink(interchaintest.InterchainLink{
			Chain1:  migaloo,
			Chain2:  migaloo2,
			Relayer: r,
			Path:    path,
		})

	ctx := context.Background()

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
		SkipPathCreation:  false,
	}))
	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Create some user accounts on both chains
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, migaloo, migaloo2)

	// Wait a few blocks for relayer to start and for user accounts to be created
	err = testutil.WaitForBlocks(ctx, 5, migaloo, migaloo2)
	require.NoError(t, err)

	// Get our Bech32 encoded user addresses
	migalooUser, migaloo2User := users[0], users[1]

	migalooUserAddr := migalooUser.FormattedAddress()
	// migaloo2UserAddr := migaloo2User.FormattedAddress()

	channel, err := ibc.GetTransferChannel(ctx, r, eRep, migaloo.Config().ChainID, migaloo2.Config().ChainID)
	require.NoError(t, err)

	_, contractAddr := helpers.SetupContract(t, ctx, migaloo2, migaloo2User.KeyName(), "bytecode/counter.wasm", `{"count":0}`)

	// do an ibc transfer through the memo to the other chain.
	transfer := ibc.WalletAmount{
		Address: contractAddr,
		Denom:   migaloo.Config().Denom,
		Amount:  int64(1),
	}

	memo := ibc.TransferOptions{
		Memo: fmt.Sprintf(`{"wasm":{"contract":"%s","msg":%s}}`, contractAddr, `{"increment":{}}`),
	}

	// Initial transfer. Account is created by the wasm execute is not so we must do this twice to properly set up
	transferTx, err := migaloo.SendIBCTransfer(ctx, channel.ChannelID, migalooUser.KeyName(), transfer, memo)
	require.NoError(t, err)
	migalooHeight, err := migaloo.Height(ctx)
	require.NoError(t, err)

	// TODO: Remove when the relayer is fixed
	r.Flush(ctx, eRep, path, channel.ChannelID)
	_, err = testutil.PollForAck(ctx, migaloo, migalooHeight-5, migalooHeight+25, transferTx.Packet)
	require.NoError(t, err)

	// Second time, this will make the counter == 1 since the account is now created.
	transferTx, err = migaloo.SendIBCTransfer(ctx, channel.ChannelID, migalooUser.KeyName(), transfer, memo)
	require.NoError(t, err)
	migalooHeight, err = migaloo.Height(ctx)
	require.NoError(t, err)

	// TODO: Remove when the relayer is fixed
	r.Flush(ctx, eRep, path, channel.ChannelID)
	_, err = testutil.PollForAck(ctx, migaloo, migalooHeight-5, migalooHeight+25, transferTx.Packet)
	require.NoError(t, err)

	// Get the address on the other chain's side
	addr := helpers.GetIBCHooksUserAddress(t, ctx, migaloo, channel.ChannelID, migalooUserAddr)
	require.NotEmpty(t, addr)

	// Get funds on the receiving chain
	funds := helpers.GetIBCHookTotalFunds(t, ctx, migaloo2, contractAddr, addr)
	require.Equal(t, int(1), len(funds.Data.TotalFunds))

	var ibcDenom string
	for _, coin := range funds.Data.TotalFunds {
		if strings.HasPrefix(coin.Denom, "ibc/") {
			ibcDenom = coin.Denom
			break
		}
	}
	require.NotEmpty(t, ibcDenom)

	// ensure the count also increased to 1 as expected.
	count := helpers.GetIBCHookCount(t, ctx, migaloo2, contractAddr, addr)
	require.Equal(t, int64(1), count.Data.Count)
}
