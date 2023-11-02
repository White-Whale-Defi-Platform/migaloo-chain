package v3

import (
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Migaloo v3.0.2 upgrade.
// this upgrade includes the fix for pfm
const UpgradeName = "v3.0.2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
}
