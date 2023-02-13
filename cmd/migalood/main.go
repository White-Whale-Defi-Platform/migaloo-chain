package main

import (
	"os"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/app"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/cmd/migalood/cmd"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/app/params"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
