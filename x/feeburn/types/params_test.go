package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultParams(t *testing.T) {
	p := DefaultParams()
	require.EqualValues(t, p.TxFeeBurnPercent, "0")
}

func TestValidateParams(t *testing.T) {
	tests := map[string]struct {
		percent   interface{}
		expectErr bool
	}{
		"DafaultParams, pass": {
			DefaultParams().TxFeeBurnPercent,
			false,
		},
		"lower boundary testing, pass": {
			"0",
			false,
		},
		"upper boundary testing, pass": {
			"100",
			false,
		},
		"greater 100%, fail": {
			"101",
			true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateTxFeeBurnPercent(test.percent)
			if test.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}

}
