package ante_test

import (
	"fmt"
	"math"
	"strconv"

	config "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app/params"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/x/feeburn/ante"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (suite *AnteTestSuite) TestFeeBurnDecorator() {
	listTxFeeBurnPercent := []string{"0", "10", "40", "50", "80", "100"}
	for _, percent := range listTxFeeBurnPercent {
		fmt.Println("tx fee burn percent", percent)
		suite.SetupTest() // reset
		err1 := suite.App.FeeBurnKeeper.UpdateTxFeeBurnPercent(suite.Ctx, percent)
		suite.Require().NoError(err1)

		fbd := ante.NewDeductFeeDecorator(suite.App.AccountKeeper, suite.App.BankKeeper, suite.App.FeeGrantKeeper, nil,
			suite.App.FeeBurnKeeper)
		antehandler := sdk.ChainAnteDecorators(fbd)

		priv := secp256k1.GenPrivKey()
		addr := getAddr(priv)
		accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())
		privNew := secp256k1.GenPrivKey()
		addrRecv := getAddr(privNew)

		accBalance := sdk.Coins{{Denom: config.BaseDenom, Amount: sdk.NewInt(int64(math.Pow10(18) * 2))}}
		err := suite.FundAccount(suite.Ctx, addr, accBalance)
		suite.Require().NoError(err)

		sendAmount := sdk.NewCoin(config.BaseDenom, sdk.NewInt(10))
		amount := sdk.Coins{sendAmount}
		sendMsg := banktypes.NewMsgSend(accountAddress, addrRecv, amount)
		supplyBefore := suite.App.BankKeeper.GetSupply(s.Ctx, config.BaseDenom).Amount
		fmt.Println("supplyBefore", supplyBefore)
		txBuilder := prepareCosmosTx(priv, sendMsg)
		// turn block for validator updates
		suite.App.EndBlock(abci.RequestEndBlock{Height: suite.Ctx.BlockHeight()})
		suite.App.Commit()
		_, err = antehandler(suite.Ctx, txBuilder.GetTx(), false)
		suite.Require().NoError(err, "Did not error on invalid tx")
		supplyAfter := suite.App.BankKeeper.GetSupply(s.Ctx, config.BaseDenom).Amount
		fmt.Println("supplyAfter", supplyAfter)
		totalTxFee := txBuilder.GetTx().GetFee()[0].Amount
		txFeeBurnPercentInt, _ := strconv.Atoi(percent)
		totalFeeBurn := totalTxFee.Mul(sdk.NewInt(int64(txFeeBurnPercentInt))).Quo(sdk.NewInt(100))
		fmt.Printf("totalTxFee %v, totalFeeBurn %v\n", totalTxFee, totalFeeBurn)
		suite.Require().True(totalFeeBurn.Equal(supplyBefore.Sub(supplyAfter)))
	}
}

func (suite *AnteTestSuite) TestFeeBurnDecoratorWhenTxNull() {
	suite.SetupTest() // reset

	fbd := ante.NewDeductFeeDecorator(suite.App.AccountKeeper, suite.App.BankKeeper, suite.App.FeeGrantKeeper, nil,
		suite.App.FeeBurnKeeper)
	antehandler := sdk.ChainAnteDecorators(fbd)

	priv := secp256k1.GenPrivKey()
	addr := getAddr(priv)
	accBalance := sdk.Coins{{Denom: config.BaseDenom, Amount: sdk.NewInt(int64(math.Pow10(18) * 2))}}
	err := suite.FundAccount(suite.Ctx, addr, accBalance)
	suite.Require().NoError(err)
	_, err = antehandler(suite.Ctx, nil, false)
	suite.Require().Error(err, "Tx must be a FeeTx")
}
