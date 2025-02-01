package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	"cosmossdk.io/log"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	customtypes "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/x/bank/keeper"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/x/tokenfactory/types"
)

type (
	Keeper struct {
		storeKey  storetypes.StoreKey
		permAddrs map[string][]string

		accountKeeper  types.AccountKeeper
		bankKeeper     *customtypes.Keeper
		contractKeeper types.ContractKeeper

		communityPoolKeeper types.CommunityPoolKeeper

		cdc       codec.BinaryCodec
		authority string
	}
)

// NewKeeper returns a new instance of the x/tokenfactory keeper
func NewKeeper(
	storeKey storetypes.StoreKey,
	permAddrs map[string][]string,
	accountKeeper types.AccountKeeper,
	bankKeeper *customtypes.Keeper,
	communityPoolKeeper types.CommunityPoolKeeper,
	cdc codec.BinaryCodec,
	authority string,
) Keeper {

	return Keeper{
		storeKey:  storeKey,
		permAddrs: permAddrs,

		accountKeeper:       accountKeeper,
		bankKeeper:          bankKeeper,
		communityPoolKeeper: communityPoolKeeper,

		cdc:       cdc,
		authority: authority,
	}
}

// Logger returns a logger for the x/tokenfactory module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetDenomPrefixStore returns the substore for a specific denom
func (k Keeper) GetDenomPrefixStore(ctx sdk.Context, denom string) storetypes.KVStore {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.GetDenomPrefixStore(denom))
}

// GetCreatorPrefixStore returns the substore for a specific creator address
func (k Keeper) GetCreatorPrefixStore(ctx sdk.Context, creator string) storetypes.KVStore {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.GetCreatorPrefix(creator))
}

// GetCreatorsPrefixStore returns the substore that contains a list of creators
func (k Keeper) GetCreatorsPrefixStore(ctx sdk.Context) storetypes.KVStore {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.GetCreatorsPrefix())
}

// Set the wasm keeper.
func (k *Keeper) SetContractKeeper(contractKeeper types.ContractKeeper) {
	k.contractKeeper = contractKeeper
}

// CreateModuleAccount creates a module account with minting and burning capabilities
// This account isn't intended to store any coins,
// it purely mints and burns them on behalf of the admin of respective denoms,
// and sends to the relevant address.
func (k Keeper) CreateModuleAccount(ctx sdk.Context) {
	k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
}

func (k Keeper) GetAuthority() string {
	return k.authority
}
