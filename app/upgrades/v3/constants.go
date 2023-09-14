package v3

import (
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	ibchookstypes "github.com/terra-money/core/v2/x/ibc-hooks/types"
)

// UpgradeName defines the on-chain upgrade name for the Migaloo v2 upgrade.
const UpgradeName = "v2.2.5"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{ibchookstypes.StoreKey},
	},
}
