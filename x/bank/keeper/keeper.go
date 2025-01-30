package keeper

import (
	"context"

	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	accountkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	customterratypes "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/x/bank/types"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/x/feeburn/types"
	custombankkeeper "github.com/terra-money/alliance/custom/bank/keeper"
)

type Keeper struct {
	custombankkeeper.Keeper
	hooks customterratypes.BankHooks
	ak    accountkeeper.AccountKeeper
}

var _ bankkeeper.Keeper = Keeper{}

func NewBaseKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	ak accountkeeper.AccountKeeper,
	blockedAddrs map[string]bool,
	authority string,
	logger log.Logger,
) Keeper {
	// add the module name to the logger
	logger = logger.With(log.ModuleKey, "x/"+types.ModuleName)

	keeper := Keeper{
		Keeper: custombankkeeper.NewBaseKeeper(cdc, storeService, ak, blockedAddrs, authority, logger),
		hooks:  nil,
		ak:     ak,
	}

	return keeper
}

// Set the bank hooks
func (k *Keeper) SetHooks(bh customterratypes.BankHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set bank hooks twice")
	}

	k.hooks = bh

	return k
}

// SendCoins transfers amt coins from a sending account to a receiving account.
// An error is returned upon failure.
func (k Keeper) SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	err := k.BlockBeforeSend(sdkctx, fromAddr, toAddr, amt)
	if err != nil {
		return err
	}
	k.TrackBeforeSend(sdkctx, fromAddr, toAddr, amt)

	return k.Keeper.SendCoins(ctx, fromAddr, toAddr, amt)
}

// SendCoinsFromAccountToModule transfers coins from an AccAddress to a ModuleAccount.
// It will panic if the module account does not exist.
func (k Keeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	recipientAcc := k.ak.GetModuleAccount(ctx, recipientModule)
	if recipientAcc == nil {
		panic(errorsmod.Wrapf(customterratypes.ErrUnknownAddress, "module account %s does not exist", recipientModule))
	}
	err := k.BlockBeforeSend(sdkctx, senderAddr, recipientAcc.GetAddress(), amt)
	if err != nil {
		return err
	}
	k.TrackBeforeSend(sdkctx, senderAddr, recipientAcc.GetAddress(), amt)

	return k.Keeper.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount transfers coins from a ModuleAccount to an AccAddress.
// It will panic if the module account does not exist. An error is returned if
// the recipient address is black-listed or if sending the tokens fails.
func (k Keeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	senderAddr := k.ak.GetModuleAddress(senderModule)
	if senderAddr == nil {
		panic(errorsmod.Wrapf(customterratypes.ErrUnknownAddress, "module account %s does not exist", senderModule))
	}
	err := k.BlockBeforeSend(sdkctx, senderAddr, recipientAddr, amt)
	if err != nil {
		return err
	}
	k.TrackBeforeSend(sdkctx, senderAddr, recipientAddr, amt)

	return k.Keeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}

// SendCoinsFromModuleToManyAccounts transfers coins from a ModuleAccount to multiple AccAddresses.
// It will panic if the module account does not exist. An error is returned if
// the recipient address is black-listed or if sending the tokens fails.
func (k Keeper) SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt sdk.Coins) error {
	senderAddr := k.ak.GetModuleAddress(senderModule)
	if senderAddr == nil {
		panic(errorsmod.Wrapf(customterratypes.ErrUnknownAddress, "senderModule address %s is nil", senderModule))
	}
	recipientAcc := k.ak.GetModuleAccount(ctx, recipientModule)
	if recipientAcc == nil {
		panic(errorsmod.Wrapf(customterratypes.ErrUnknownAddress, "recipientModule address %s is nil", recipientModule))
	}

	return k.Keeper.SendCoins(ctx, senderAddr, recipientAcc.GetAddress(), amt)
}

// UndelegateCoins performs undelegation by crediting amt coins to an account with
// address addr. For vesting accounts, undelegation amounts are tracked for both
// vesting and vested coins. The coins are then transferred from a ModuleAccount
// address to the delegator address. If any of the undelegation amounts are
// negative, an error is returned.
func (k Keeper) UndelegateCoins(ctx context.Context, moduleAccAddr, delegatorAddr sdk.AccAddress, amt sdk.Coins) error {
	sdkctx := sdk.UnwrapSDKContext(ctx)

	err := k.BlockBeforeSend(sdkctx, moduleAccAddr, delegatorAddr, amt)
	if err != nil {
		return err
	}
	k.TrackBeforeSend(sdkctx, moduleAccAddr, delegatorAddr, amt)

	return k.Keeper.UndelegateCoins(ctx, moduleAccAddr, delegatorAddr, amt)
}

// DelegateCoins performs delegation by deducting amt coins from an account with
// address addr. For vesting accounts, delegations amounts are tracked for both
// vesting and vested coins. The coins are then transferred from the delegator
// address to a ModuleAccount address. If any of the delegation amounts are negative,
// an error is returned.
func (k Keeper) DelegateCoins(ctx context.Context, delegatorAddr, moduleAccAddr sdk.AccAddress, amt sdk.Coins) error {
	sdkctx := sdk.UnwrapSDKContext(ctx)

	err := k.BlockBeforeSend(sdkctx, delegatorAddr, moduleAccAddr, amt)
	if err != nil {
		return err
	}
	k.TrackBeforeSend(sdkctx, delegatorAddr, moduleAccAddr, amt)

	return k.Keeper.DelegateCoins(ctx, delegatorAddr, moduleAccAddr, amt)
}

// InputOutputCoins performs multi-send functionality. It accepts a series of
// input that correspond to a series of outputs. It returns an error if the
// input and outputs don't line up or if any single transfer of tokens fails.
func (k Keeper) InputOutputCoins(ctx context.Context, input banktypes.Input, outputs []banktypes.Output) error {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	// Safety check ensuring that when sending coins the keeper must maintain the
	// Check supply invariant and validity of Coins.
	if err := banktypes.ValidateInputOutputs(input, outputs); err != nil {
		return err
	}

	inputaddress := sdk.MustAccAddressFromBech32(input.Address)

	for _, output := range outputs {
		outputaddress := sdk.MustAccAddressFromBech32(output.Address)

		err := k.BlockBeforeSend(sdkctx, inputaddress, outputaddress, output.Coins)
		if err != nil {
			return err
		}
		k.TrackBeforeSend(sdkctx, inputaddress, outputaddress, output.Coins)
	}

	return k.Keeper.InputOutputCoins(ctx, input, outputs)
}
