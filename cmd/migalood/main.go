package main

import (
	"os"

	"cosmossdk.io/log"
	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/cmd/migalood/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "MIGALOOD", app.DefaultNodeHome); err != nil {
		log.NewLogger(rootCmd.OutOrStderr()).Error("failure when running app", "err", err)
		os.Exit(1)
	}
}
