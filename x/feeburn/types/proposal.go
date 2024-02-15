package types

import (
	"fmt"
	"strconv"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeUpdateTxFeeBurnPercent string = "msg_update_tx_fee_burn_percent"
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &MsgUpdateTxFeeBurnPercentProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeUpdateTxFeeBurnPercent)
}

func NewMsgUpdateTxFeeBurnPercentProposal(title, description, txFeeBurnPercent string) govtypes.Content {
	return &MsgUpdateTxFeeBurnPercentProposal{
		Title:            title,
		Description:      description,
		TxFeeBurnPercent: txFeeBurnPercent,
	}
}

func (m *MsgUpdateTxFeeBurnPercentProposal) GetTitle() string       { return m.Title }
func (m *MsgUpdateTxFeeBurnPercentProposal) GetDescription() string { return m.Description }

// ProposalRoute returns router key for this proposal
func (m *MsgUpdateTxFeeBurnPercentProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (m *MsgUpdateTxFeeBurnPercentProposal) ProposalType() string {
	return ProposalTypeUpdateTxFeeBurnPercent
}

// ValidateBasic performs a stateless check of the proposal fields
func (m *MsgUpdateTxFeeBurnPercentProposal) ValidateBasic() error {
	txFeeBurnPercentInt, err := strconv.Atoi(m.TxFeeBurnPercent)
	if err != nil {
		return err
	}
	if txFeeBurnPercentInt < 0 || txFeeBurnPercentInt > 100 {
		return fmt.Errorf("fee must be between 0 and 100")
	}
	return govtypes.ValidateAbstract(m)
}
