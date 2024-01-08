package keeper

import (
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UpdateTxFeeBurnPercent sets a new transaction fee burn percentage after validation.
func (k Keeper) UpdateTxFeeBurnPercent(ctx sdk.Context, newTxFeeBurnPercent string) error {
	ms := MsgServer{k}

	newParams := types.NewParams(newTxFeeBurnPercent)
	if err := newParams.Validate(); err != nil {
		return err
	}
	_, err := ms.UpdateParams(ctx, &types.MsgUpdateParams{
		Authority: k.GetAuthority(),
		Params:    newParams,
	})
	return err

}
