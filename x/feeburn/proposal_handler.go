package feeburn

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/x/feeburn/keeper"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/x/feeburn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// NewFeeBurnProposalHandler returns a handler for FeeBurn proposals.
func NewFeeBurnProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.MsgUpdateTxFeeBurnPercentProposal:
			return handleUpdateTxFeeBurnPercentProposal(ctx, k, c)

		default:
			return errorsmod.Wrapf(errortypes.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleUpdateTxFeeBurnPercentProposal(
	ctx sdk.Context,
	k keeper.Keeper,
	p *types.MsgUpdateTxFeeBurnPercentProposal,
) error {
	err := k.UpdateTxFeeBurnPercent(ctx, p.TxFeeBurnPercent)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateTxBurnFeePercent,
			sdk.NewAttribute(types.AttributeKeyTxBurnFeePercent, p.TxFeeBurnPercent),
		),
	)

	return nil
}
