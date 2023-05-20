package interchaintest

import (
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
)

var (
	MigalooMainRepo = "ghcr.io/white-whale-defi-platform/migaloo-chain"

	MigalooImage = ibc.DockerImage{
		Repository: "ghcr.io/white-whale-defi-platform/migaloo-chain-ictest",
		Version:    "latest",
		UidGid:     "1025:1025",
	}

	migalooConfig = ibc.ChainConfig{
		Type:                "cosmos",
		Name:                "migaloo",
		ChainID:             "migaloo-2",
		Images:              []ibc.DockerImage{MigalooImage},
		Bin:                 "migalood",
		Bech32Prefix:        "migaloo",
		Denom:               "stake",
		CoinType:            "118",
		GasPrices:           "0.0stake",
		GasAdjustment:       1.1,
		TrustingPeriod:      "112h",
		NoHostMount:         false,
		ModifyGenesis:       nil,
		ConfigFileOverrides: nil,
	}
)
