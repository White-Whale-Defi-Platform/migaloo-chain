package keeper_test

import (
	"testing"

	testkeeper "github.com/White-Whale-Defi-Platform/migaloo-chain/v3/testutil/keeper"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.FeeburnKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	err := keeper.SetParams(ctx, params)
	require.NoError(t, err)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
