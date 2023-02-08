package params

// Default simulation operation weights for messages and gov proposals
const (
	DefaultWeightMsgSend                        int = 100
	DefaultWeightMsgMultiSend                   int = 10
	DefaultWeightMsgSetWithdrawAddress          int = 50
	DefaultWeightMsgWithdrawDelegationReward    int = 50
	DefaultWeightMsgWithdrawValidatorCommission int = 50
	DefaultWeightMsgFundCommunityPool           int = 50
	DefaultWeightMsgDeposit                     int = 100
	DefaultWeightMsgVote                        int = 67
	DefaultWeightMsgUnjail                      int = 100
	DefaultWeightMsgCreateValidator             int = 100
	DefaultWeightMsgEditValidator               int = 5
	DefaultWeightMsgDelegate                    int = 100
	DefaultWeightMsgUndelegate                  int = 100
	DefaultWeightMsgBeginRedelegate             int = 100

	DefaultWeightCommunitySpendProposal int = 5
	DefaultWeightTextProposal           int = 5
	DefaultWeightParamChangeProposal    int = 5

	DefaultWeightMsgStoreCode           int = 50
	DefaultWeightMsgInstantiateContract int = 100
	DefaultWeightMsgExecuteContract     int = 100
	DefaultWeightMsgUpdateAdmin         int = 25
	DefaultWeightMsgClearAdmin          int = 10
	DefaultWeightMsgMigrateContract     int = 50

	DefaultWeightStoreCodeProposal                   int = 5
	DefaultWeightInstantiateContractProposal         int = 5
	DefaultWeightUpdateAdminProposal                 int = 5
	DefaultWeightExecuteContractProposal             int = 5
	DefaultWeightClearAdminProposal                  int = 5
	DefaultWeightMigrateContractProposal             int = 5
	DefaultWeightSudoContractProposal                int = 5
	DefaultWeightPinCodesProposal                    int = 5
	DefaultWeightUnpinCodesProposal                  int = 5
	DefaultWeightUpdateInstantiateConfigProposal     int = 5
	DefaultWeightStoreAndInstantiateContractProposal int = 5

	//token factory
	DefaultWeightMsgCreateDenom        int = 100
	DefaultWeightMsgMint               int = 100
	DefaultWeightMsgBurn               int = 100
	DefaultWeightMsgEditDenom          int = 100
	DefaultWeightMsgTransferDenomOwner int = 100
	DefaultWeightMsgChangeAdmin        int = 100
	DefaultWeightMsgSetDenomMetadata   int = 100

	DefaultWeightMsgAddDenomMetadata                int = 100
	DefaultWeightMsgRemoveDenomMetadata             int = 100
	DefaultWeightMsgAddDenomMetadataAddress         int = 100
	DefaultWeightMsgRemoveDenomMetadataAddress      int = 100
	DefaultWeightMsgAddDenomMetadataURI             int = 100
	DefaultWeightMsgRemoveDenomMetadataURI          int = 100
	DefaultWeightMsgAddDenomMetadataHash            int = 100
	DefaultWeightMsgRemoveDenomMetadataHash         int = 100
	DefaultWeightMsgAddDenomMetadataDescription     int = 100
	DefaultWeightMsgRemoveDenomMetadataDescription  int = 100
	DefaultWeightMsgAddDenomMetadataName            int = 100
	DefaultWeightMsgRemoveDenomMetadataName         int = 100
	DefaultWeightMsgAddDenomMetadataSymbol          int = 100
	DefaultWeightMsgRemoveDenomMetadataSymbol       int = 100
	DefaultWeightMsgAddDenomMetadataDecimals        int = 100
	DefaultWeightMsgRemoveDenomMetadataDecimals     int = 100
	DefaultWeightMsgAddDenomMetadataTotalSupply     int = 100
	DefaultWeightMsgRemoveDenomMetadataTotalSupply  int = 100
	DefaultWeightMsgAddDenomMetadataMintable        int = 100
	DefaultWeightMsgRemoveDenomMetadataMintable     int = 100
	DefaultWeightMsgAddDenomMetadataBurnable        int = 100
	DefaultWeightMsgRemoveDenomMetadataBurnable     int = 100
	DefaultWeightMsgAddDenomMetadataTransferable    int = 100
	DefaultWeightMsgRemoveDenomMetadataTransferable int = 100
	DefaultWeightMsgAddDenomMetadataRestricted      int = 100
	DefaultWeightMsgRemoveDenomMetadataRestricted   int = 100
	DefaultWeightMsgAddDenomMetadataMaxSupply       int = 100
	DefaultWeightMsgRemoveDenomMetadataMaxSupply    int = 100
	DefaultWeightMsgAddDenomMetadataDenomUnits      int = 100
	DefaultWeightMsgRemoveDenomMetadataDenomUnits   int = 100
)
