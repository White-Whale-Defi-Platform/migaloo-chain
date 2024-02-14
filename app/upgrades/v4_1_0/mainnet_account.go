package v4

import (
	"encoding/json"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"
)

func CreateMainnetVestingAccount(ctx sdk.Context,
	bankKeeper bankKeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
) (vestingtypes.ContinuousVestingAccount, math.Int) {
	str := `{"@type":"/cosmos.vesting.v1beta1.ContinuousVestingAccount","base_vesting_account":{"base_account":{"address":"migaloo1alga5e8vr6ccr9yrg0kgxevpt5xgmgrvqgujs6","pub_key":{"@type":"/cosmos.crypto.multisig.LegacyAminoPubKey","threshold":4,"public_keys":[{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AlnzK22KrkylnvTCvZZc8eZnydtQuzCWLjJJSMFUvVHf"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aiw2Ftg+fnoHDU7M3b0VMRsI0qurXlerW0ahtfzSDZA4"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AvEHv+MVYRVau8FbBcJyG0ql85Tbbn7yhSA0VGmAY4ku"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Az5VHWqi3zMJu1rLGcu2EgNXLLN+al4Dy/lj6UZTzTCl"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Ai4GlSH3uG+joMnAFbQC3jQeHl9FPvVTlRmwIFt7d7TI"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2kAzH2bZr530jmFq/bRFrT2q8SRqdnfIebba+YIBqI1"}]},"account_number":46,"sequence":27},"original_vesting":[{"denom":"uwhale","amount":"22165200000000"}],"delegated_free":[{"denom":"uwhale","amount":"443382497453"}],"delegated_vesting":[{"denom":"uwhale","amount":"22129422502547"}],"end_time":1770994800},"start_time":1676300400}`

	var acc vestingtypes.ContinuousVestingAccount
	if err := json.Unmarshal([]byte(str), &acc); err != nil {
		panic(err)
	}

	vesting := GetVestingCoin(ctx, &acc)

	err := banktestutil.FundAccount(bankKeeper, ctx, acc.BaseAccount.GetAddress(),
		acc.GetOriginalVesting())
	if err != nil {
		panic(err)
	}

	accountKeeper.SetAccount(ctx, &acc)
	return acc, vesting
}
