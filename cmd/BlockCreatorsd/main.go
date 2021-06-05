package main

import (
	"os"

	"github.com/BlockCreatorsNetwork/BlockCreators/app"
	"github.com/BlockCreatorsNetwork/BlockCreators/cmd/BlockCreatorsd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
