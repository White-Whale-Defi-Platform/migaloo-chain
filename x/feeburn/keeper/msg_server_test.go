package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/White-Whale-Defi-Platform/migaloo-chain/v3/testutil/keeper"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/keeper"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FeeburnKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
