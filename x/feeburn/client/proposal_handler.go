package client

import (
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/client/cli"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var UpdateTxFeeBurnPercentProposalHandler = govclient.NewProposalHandler(cli.NewUpdateTxFeeBurnPercentProposalHandler)
