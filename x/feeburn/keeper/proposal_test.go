package keeper_test

func (suite *KeeperTestSuite) TestUpdateTxFeeBurnPercent() {
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
			name:    "invalid fee",
			fee:     "abc",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := suite.app.FeeBurnKeeper.UpdateTxFeeBurnPercent(suite.ctx, tt.fee)
			if tt.wantErr {
				suite.NotNil(err)
			} else {
				suite.Nil(err)
				params := suite.app.FeeBurnKeeper.GetParams(suite.ctx)
				suite.Equal(tt.fee, params.TxFeeBurnPercent)
			}
		})
	}
}
