package v4

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	alliancekeeper "github.com/terra-money/alliance/x/alliance/keeper"
)

func UpdateAlliance(ctx sdk.Context, alk alliancekeeper.Keeper) {
	allianceParams := alk.GetParams(ctx)
	allianceParams.RewardDelayTime = 10 * time.Minute

	_ = alk.SetParams(ctx, allianceParams)

	asset, found := alk.GetAssetByDenom(ctx, "ibc/30E9709461C4DA26A7A579E11DE44B591E676DB1B7F94714FBFF87ED6E47D6F4")
	if found {
		rewardStartTime := ctx.BlockTime()
		asset.LastRewardChangeTime = rewardStartTime
		asset.RewardStartTime = rewardStartTime
		alk.SetAsset(ctx, asset)
	}
}
