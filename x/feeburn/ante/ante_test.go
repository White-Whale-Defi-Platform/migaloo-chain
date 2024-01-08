package ante_test

import (
	config "github.com/White-Whale-Defi-Platform/migaloo-chain/v3/app/params"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/ante"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"math"
)

func (suite *AnteTestSuite) TestFeeBurnDecorator() {
	suite.SetupTest(false) // reset

	fbd := ante.NewDeductFeeDecorator(suite.App.AccountKeeper, suite.App.BankKeeper, suite.App.FeeGrantKeeper, nil,
		suite.App.FeeBurnKeeper)
	antehandler := sdk.ChainAnteDecorators(fbd)

	priv := ed25519.GenPrivKey()
	addr := getAddr(priv)
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())
	privNew := ed25519.GenPrivKey()
	addrRecv := getAddr(privNew)

	accBalance := sdk.Coins{{Denom: config.BaseDenom, Amount: sdk.NewInt(int64(math.Pow10(18) * 2))}}
	err := suite.FundAccount(suite.Ctx, addr, accBalance)
	suite.Require().NoError(err)

	sendAmount := sdk.NewCoin(config.BaseDenom, sdk.NewInt(10))
	amount := sdk.Coins{sendAmount}
	sendMsg := banktypes.NewMsgSend(accountAddress, addrRecv, amount)
	txBuilder := prepareCosmosTx(priv, sendMsg)
	_, err = antehandler(suite.Ctx, txBuilder.GetTx(), false)

	suite.Require().NoError(err, "Did not error on invalid tx")
}

func (suite *AnteTestSuite) TestFeeBurnDecoratorWhenTxNull() {
	suite.SetupTest(false) // reset

	fbd := ante.NewDeductFeeDecorator(suite.App.AccountKeeper, suite.App.BankKeeper, suite.App.FeeGrantKeeper, nil,
		suite.App.FeeBurnKeeper)
	antehandler := sdk.ChainAnteDecorators(fbd)

	priv := ed25519.GenPrivKey()
	addr := getAddr(priv)
	accBalance := sdk.Coins{{Denom: config.BaseDenom, Amount: sdk.NewInt(int64(math.Pow10(18) * 2))}}
	err := suite.FundAccount(suite.Ctx, addr, accBalance)
	suite.Require().NoError(err)
	_, err = antehandler(suite.Ctx, nil, false)
	suite.Require().Error(err, "Tx must be a FeeTx")
}
