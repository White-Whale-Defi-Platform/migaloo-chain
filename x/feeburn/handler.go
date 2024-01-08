package feeburn

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for messages in the FeeBurn module.
// It handles the MsgUpdateParams type and returns an error for unrecognized message types.
func NewHandler(server types.MsgServer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (result *sdk.Result, err error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgUpdateParams:
			res, err := server.UpdateParams(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			err := sdkerrors.Wrapf(errortypes.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, err
		}
	}
}
