package v2

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Migaloo v2 upgrade.
const UpgradeName = "v2.2.5"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{},
	},
}
