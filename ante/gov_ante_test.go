package ante_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/stretchr/testify/suite"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/ante"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	migalooapp "github.com/White-Whale-Defi-Platform/migaloo-chain/v3/app"
)

var (
	insufficientCoins = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 100))
	minCoins          = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000))
	moreThanMinCoins  = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 2500000))
	testAddr          = sdk.AccAddress("test1")
)

type GovAnteHandlerTestSuite struct {
	suite.Suite

	app       *migalooapp.MigalooApp
	ctx       sdk.Context
	clientCtx client.Context
}

func (s *GovAnteHandlerTestSuite) SetupTest() {
	app := migalooapp.SetupMigalooAppWithValSet(s.T())
	ctx := app.BaseApp.NewContext(false, tmproto.Header{
		ChainID: fmt.Sprintf("test-chain-%s", tmrand.Str(4)),
		Height:  1,
	})

	encodingConfig := migalooapp.MakeEncodingConfig()
	encodingConfig.Amino.RegisterConcrete(&testdata.TestMsg{}, "testdata.TestMsg", nil)
	testdata.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	s.app = app
	s.ctx = ctx
	s.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
}

func TestGovSpamPreventionSuite(t *testing.T) {
	suite.Run(t, new(GovAnteHandlerTestSuite))
}

func (s *GovAnteHandlerTestSuite) TestGlobalFeeMinimumGasFeeAnteHandler() {
	// setup test
	s.SetupTest()
	tests := []struct {
		title, description string
		proposalType       string
		proposerAddr       sdk.AccAddress
		initialDeposit     sdk.Coins
		expectPass         bool
	}{
		{"Passing proposal 1", "the purpose of this proposal is to pass", govv1beta1.ProposalTypeText, testAddr, minCoins, true},
		{"Passing proposal 2", "the purpose of this proposal is to pass with more coins than minimum", govv1beta1.ProposalTypeText, testAddr, moreThanMinCoins, true},
		{"Failing proposal", "the purpose of this proposal is to fail", govv1beta1.ProposalTypeText, testAddr, insufficientCoins, false},
	}

	decorator := ante.NewGovPreventSpamDecorator(s.app.AppCodec(), &s.app.GovKeeper)

	for _, tc := range tests {
		content, ok := govv1beta1.ContentFromProposalType(tc.title, tc.description, tc.proposalType)
		s.Require().True(ok)
		s.Require().NotNil(content)
		msg, err := govv1beta1.NewMsgSubmitProposal(
			content,
			tc.initialDeposit,
			tc.proposerAddr,
		)

		s.Require().NoError(err)

		err = decorator.ValidateGovMsgs(s.ctx, []sdk.Msg{msg})
		if tc.expectPass {
			s.Require().NoError(err, "expected %v to pass", tc.title)
		} else {
			s.Require().Error(err, "expected %v to fail", tc.title)
		}
	}
}
