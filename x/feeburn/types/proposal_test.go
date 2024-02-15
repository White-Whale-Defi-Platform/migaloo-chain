package types

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProposalTestSuite struct {
	suite.Suite
}

func TestProposalTestSuite(t *testing.T) {
	suite.Run(t, new(ProposalTestSuite))
}

func (suite *ProposalTestSuite) TestNewUpdateTxFeeBurnPercentProposal() {
	title := "Test Title"
	description := "Test Description"
	txFeeBurnPercent := "50"

	proposal := NewMsgUpdateTxFeeBurnPercentProposal(title, description, txFeeBurnPercent)
	err := proposal.ValidateBasic()
	suite.Require().NoError(err)

	suite.Require().Equal(title, proposal.(*MsgUpdateTxFeeBurnPercentProposal).Title)
	suite.Require().Equal(txFeeBurnPercent, proposal.(*MsgUpdateTxFeeBurnPercentProposal).TxFeeBurnPercent)
	suite.Require().Equal(description, proposal.(*MsgUpdateTxFeeBurnPercentProposal).Description)
}

func (suite *ProposalTestSuite) TestUpdateTxFeeBurnPercentProposal_ValidateBasic() {
	tests := []struct {
		name    string
		fee     string
		wantErr bool
	}{
		{
			name:    "valid fee",
			fee:     "50",
			wantErr: false,
		},
		{
			name:    "fee too high",
			fee:     "101",
			wantErr: true,
		},
		{
			name:    "fee too low",
			fee:     "-1",
			wantErr: true,
		},
		{
			name:    "non-numeric fee",
			fee:     "abc",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		u := &MsgUpdateTxFeeBurnPercentProposal{
			Title:            "Test",
			Description:      "Test Description",
			TxFeeBurnPercent: tc.fee,
		}
		err := u.ValidateBasic()
		if tc.wantErr {
			suite.Require().Error(err, tc.name)
		} else {
			suite.Require().NoError(err)
		}

	}
}
