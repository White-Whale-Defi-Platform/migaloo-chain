package keeper_test

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (suite *KeeperTestSuite) TestBurnFeeCosmosTxDelegate() {
	suite.SetupTest()
	priv0 := secp256k1.GenPrivKey()
	addr := sdk.AccAddress(priv0.PubKey().Address())
	accBalance := sdk.Coins{{Denom: sdk.DefaultBondDenom, Amount: sdkmath.NewInt(10000000000000)}}
	err := suite.FundAccount(suite.Ctx, addr, accBalance)
	suite.Require().NoError(err)
	totalSupplyBefore := suite.App.BankKeeper.GetSupply(suite.Ctx, sdk.DefaultBondDenom)
	fmt.Println("totalSupply", totalSupplyBefore)
	delegateAmount := sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(1000))
	delegate(priv0, delegateAmount)
	mintedCoin, err := getMintedCoin()
	suite.Require().NoError(err)
	totalSupplyAfter := suite.App.BankKeeper.GetSupply(suite.Ctx, sdk.DefaultBondDenom)
	fmt.Println("totalSupplyAfter", totalSupplyAfter)
	expectAmount := totalSupplyAfter.Amount.Sub(totalSupplyBefore.Amount)
	expectAmount = mintedCoin.Amount.Sub(expectAmount)
	fmt.Println("expectAmount", expectAmount)
}

func delegate(priv cryptotypes.PrivKey, delegateAmount sdk.Coin) {
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())
	validators, err := s.App.StakingKeeper.GetValidators(s.Ctx, 1)
	s.Require().NoError(err)

	val, err := sdk.ValAddressFromBech32(validators[0].OperatorAddress)
	s.Require().NoError(err)

	delegateMsg := stakingtypes.NewMsgDelegate(accountAddress.String(), val.String(), delegateAmount)
	res, err := deliverTx(priv, delegateMsg)
	s.Require().NoError(err)
	s.Require().Equal(uint32(0), res.Code)
}

func deliverTx(priv cryptotypes.PrivKey, msgs ...sdk.Msg) (*abci.ExecTxResult, error) {
	bz := prepareCosmosTx(priv, msgs...)
	res, err := s.App.FinalizeBlock(&abci.RequestFinalizeBlock{
		Txs:    [][]byte{bz},
		Height: s.Ctx.BlockHeight(),
	})

	s.Require().Len(res.TxResults, 1)
	return res.TxResults[0], err
}

func getMintedCoin() (sdk.Coin, error) {
	mintParams, err := s.App.MintKeeper.Params.Get(s.Ctx)
	if err != nil {
		return sdk.Coin{}, err
	}

	minter, err := s.App.MintKeeper.Minter.Get(s.Ctx)
	if err != nil {
		return sdk.Coin{}, err
	}

	return minter.BlockProvision(mintParams), nil
}
