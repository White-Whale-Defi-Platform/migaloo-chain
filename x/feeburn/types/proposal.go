package types

import (
	"fmt"
	govcdc "github.com/cosmos/cosmos-sdk/x/gov/codec"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"strconv"
)

const (
	ProposalTypeUpdateTxFeeBurnPercent string = "msg_update_tx_fee_burn_percent"
)

// Implements Proposal Interface
var (
	_ v1beta1.Content = &MsgUpdateTxFeeBurnPercentProposal{}
)

func init() {
	v1beta1.RegisterProposalType(ProposalTypeUpdateTxFeeBurnPercent)
	govcdc.ModuleCdc.Amino.RegisterConcrete(&MsgUpdateTxFeeBurnPercentProposal{}, "feeburn/UpdateTxFeeBurnPercentProposal", nil)
}

func NewMsgUpdateTxFeeBurnPercentProposal(title, description, txFeeBurnPercent string) v1beta1.Content {
	return &MsgUpdateTxFeeBurnPercentProposal{
		Title:            title,
		Description:      description,
		TxFeeBurnPercent: txFeeBurnPercent,
	}
}

// ProposalRoute returns router key for this proposal
func (*MsgUpdateTxFeeBurnPercentProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*MsgUpdateTxFeeBurnPercentProposal) ProposalType() string {
	return ProposalTypeUpdateTxFeeBurnPercent
}

// ValidateBasic performs a stateless check of the proposal fields
func (u *MsgUpdateTxFeeBurnPercentProposal) ValidateBasic() error {
	txFeeBurnPercentInt, err := strconv.Atoi(u.TxFeeBurnPercent)
	if err != nil {
		return err
	}
	if txFeeBurnPercentInt < 0 || txFeeBurnPercentInt > 100 {
		return fmt.Errorf("fee must be between 0 and 100")
	}
	return v1beta1.ValidateAbstract(u)
}
