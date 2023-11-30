package interchaintest

import (
	"os"
	"strings"

	"github.com/strangelove-ventures/interchaintest/v5/ibc"
)

var (
	MigalooICTestRepo = "ghcr.io/white-whale-defi-platform/migaloo-chain-ictest"
	MigalooMainRepo   = "ghcr.io/white-whale-defi-platform/migaloo-chain"

	IBCRelayerImage   = "ghcr.io/cosmos/relayer"
	IBCRelayerVersion = "justin-localhost-ibc"

	repo, version = GetDockerImageInfo()

	MigalooImage = ibc.DockerImage{
		Repository: repo,
		Version:    version,
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

// GetDockerImageInfo returns the appropriate repo and branch version string for integration with the CI pipeline.
// The remote runner sets the BRANCH_CI env var. If present, interchaintest will use the docker image pushed up to the repo.
// If testing locally, user should run `make docker-build-debug` and interchaintest will use the local image.
func GetDockerImageInfo() (repo, version string) {
	branchVersion, found := os.LookupEnv("BRANCH_CI")
	repo = MigalooICTestRepo
	if !found {
		// make local-image
		repo = "migaloo"
		branchVersion = "debug"
	}

	// github converts / to - for pushed docker images
	branchVersion = strings.ReplaceAll(branchVersion, "/", "-")
	return repo, branchVersion
}
