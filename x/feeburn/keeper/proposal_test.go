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
			suite.SetupTest()
			err := suite.App.FeeBurnKeeper.UpdateTxFeeBurnPercent(suite.Ctx, tt.fee)
			if tt.wantErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				params := suite.App.FeeBurnKeeper.GetParams(suite.Ctx)
				suite.Require().Equal(tt.fee, params.GetTxFeeBurnPercent())
			}
		})
	}
}
