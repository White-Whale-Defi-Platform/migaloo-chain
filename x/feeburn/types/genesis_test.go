package types_test

import (
	"testing"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v3/x/feeburn/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "invalid genesis state",
			genState: &types.GenesisState{
				Params: types.Params{
					TxFeeBurnPercent: "101",
				},
			},
			valid: false,
		},
		{
			desc: "invalid genesis state",
			genState: &types.GenesisState{
				Params: types.Params{
					TxFeeBurnPercent: "-1",
				},
			},
			valid: false,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.Params{
					TxFeeBurnPercent: "50",
				},
			},
			valid: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
