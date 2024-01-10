package main

import (
	"os"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/cmd/migalood/cmd"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "MIGALOOD", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
