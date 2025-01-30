package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var (
	ErrUnknownAddress = sdkerrors.Register(banktypes.ModuleName, 383838, "module account does not exist")
	// ErrUnauthorized is used whenever a request without sufficient
	// authorization is handled.
	ErrUnauthorized = sdkerrors.Register(banktypes.ModuleName, 383839, "unauthorized")
)
