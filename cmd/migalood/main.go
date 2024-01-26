package main

import (
	"fmt"
	"os"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/cmd/migalood/cmd"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/v4/app"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	fmt.Println("app start here")

	if err := svrcmd.Execute(rootCmd, "MIGALOOD", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
