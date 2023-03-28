package v2

import (
	"github.com/White-Whale-Defi-Platform/migaloo-chain/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	alliancetypes "github.com/terra-money/alliance/x/alliance/types"
)

// UpgradeName defines the on-chain upgrade name for the Migaloo v2 upgrade.
const UpgradeName = "v2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			alliancetypes.ModuleName,
		},
	},
}
