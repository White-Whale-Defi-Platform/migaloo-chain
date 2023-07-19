package v2_2_5

import (
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v2/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

// UpgradeName defines the on-chain upgrade name for the Migaloo v2 upgrade.
const UpgradeName = "v2.2.5"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}
