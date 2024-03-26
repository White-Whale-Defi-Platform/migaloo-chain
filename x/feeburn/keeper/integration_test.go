package keeper_test

import (
	"fmt"

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
	accBalance := sdk.Coins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewInt(10000000000000)}}
	err := suite.FundAccount(suite.Ctx, addr, accBalance)
	suite.Require().NoError(err)
	totalSupplyBefore := suite.App.BankKeeper.GetSupply(suite.Ctx, sdk.DefaultBondDenom)
	fmt.Println("totalSupply", totalSupplyBefore)
	delegateAmount := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
	delegate(priv0, delegateAmount)
	mintedCoin := getMintedCoin()
	totalSupplyAfter := suite.App.BankKeeper.GetSupply(suite.Ctx, sdk.DefaultBondDenom)
	fmt.Println("totalSupplyAfter", totalSupplyAfter)
	expectAmount := totalSupplyAfter.Amount.Sub(totalSupplyBefore.Amount)
	expectAmount = mintedCoin.Amount.Sub(expectAmount)
	fmt.Println("expectAmount", expectAmount)
}

func delegate(priv cryptotypes.PrivKey, delegateAmount sdk.Coin) {
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())
	validators := s.App.StakingKeeper.GetValidators(s.Ctx, 1)

	val, err := sdk.ValAddressFromBech32(validators[0].OperatorAddress)
	s.Require().NoError(err)

	delegateMsg := stakingtypes.NewMsgDelegate(accountAddress, val, delegateAmount)
	res := deliverTx(priv, delegateMsg)
	s.Require().Equal(uint32(0), res.Code)
}

func deliverTx(priv cryptotypes.PrivKey, msgs ...sdk.Msg) abci.ResponseDeliverTx {
	bz := prepareCosmosTx(priv, msgs...)
	req := abci.RequestDeliverTx{Tx: bz}
	res := s.App.BaseApp.DeliverTx(req)
	return res
}

func getMintedCoin() sdk.Coin {
	mintParams := s.App.MintKeeper.GetParams(s.Ctx)
	return s.App.MintKeeper.GetMinter(s.Ctx).BlockProvision(mintParams)
}
