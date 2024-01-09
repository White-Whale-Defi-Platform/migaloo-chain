package ante_test

import (
	"testing"

	apptesting "github.com/White-Whale-Defi-Platform/migaloo-chain/v3/app"
	config "github.com/White-Whale-Defi-Platform/migaloo-chain/v3/app/params"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/ibc-go/v7/testing/simapp"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

var (
	DefaultFee int64 = 1000000000000
	s          *AnteTestSuite
)

// AnteTestSuite is a test suite to be used with ante handler tests.
type AnteTestSuite struct {
	apptesting.KeeperTestHelper
	anteHandler sdk.AnteHandler
	clientCtx   client.Context
}

// SetupTest setups a new test, with new app, context, and anteHandler.
func (suite *AnteTestSuite) SetupTest() {
	suite.Setup(suite.T(), apptesting.SimAppChainID)

	// Set up TxConfig.
	encodingConfig := simapp.MakeTestEncodingConfig()
	// We're using TestMsg encoding in some tests, so register it here.
	encodingConfig.Amino.RegisterConcrete(&testdata.TestMsg{}, "testdata.TestMsg", nil)
	testdata.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	suite.clientCtx = client.Context{}.
		WithTxConfig(encodingConfig.TxConfig)

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   suite.App.AccountKeeper,
			BankKeeper:      suite.App.BankKeeper,
			FeegrantKeeper:  suite.App.FeeGrantKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)

	suite.Require().NoError(err)
	suite.anteHandler = anteHandler
}

func TestAnteTestSuite(t *testing.T) {
	s = new(AnteTestSuite)
	suite.Run(t, s)
}

func getAddr(priv cryptotypes.PrivKey) sdk.AccAddress {
	return priv.PubKey().Address().Bytes()
}

func prepareCosmosTx(priv cryptotypes.PrivKey, msgs ...sdk.Msg) client.TxBuilder {
	encodingConfig := config.MakeEncodingConfig()
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())

	txBuilder := encodingConfig.TxConfig.NewTxBuilder()

	txBuilder.SetGasLimit(1000000)
	gasPrice := sdk.NewInt(1)
	fees := &sdk.Coins{{Denom: config.BaseDenom, Amount: gasPrice.MulRaw(DefaultFee)}}
	txBuilder.SetFeeAmount(*fees)
	err := txBuilder.SetMsgs(msgs...)
	s.Require().NoError(err)

	seq, err := s.App.AccountKeeper.GetSequence(s.Ctx, accountAddress)
	s.Require().NoError(err)

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	sigV2 := signing.SignatureV2{
		PubKey: priv.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encodingConfig.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: seq,
	}

	sigsV2 := []signing.SignatureV2{sigV2}

	err = txBuilder.SetSignatures(sigsV2...)
	s.Require().NoError(err)

	// Second round: all signer infos are set, so each signer can sign.
	accNumber := s.App.AccountKeeper.GetAccount(s.Ctx, accountAddress).GetAccountNumber()
	signerData := authsigning.SignerData{
		ChainID:       s.Ctx.ChainID(),
		AccountNumber: accNumber,
		Sequence:      seq,
	}
	sigV2, err = tx.SignWithPrivKey(
		encodingConfig.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, priv, encodingConfig.TxConfig,
		seq,
	)
	s.Require().NoError(err)

	sigsV2 = []signing.SignatureV2{sigV2}
	err = txBuilder.SetSignatures(sigsV2...)
	s.Require().NoError(err)

	return txBuilder
}

func (suite *AnteTestSuite) FundAccount(ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := suite.App.BankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return suite.App.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}
