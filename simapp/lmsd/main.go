package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	// "github.com/cosmos/cosmos-sdk/simapp"
	// "lms-cosmos/simapp"
	"github.com/Leave-Management-System/lms-cosmos/simapp"

	// "github.com/cosmos/cosmos-sdk/simapp/simd/cmd"
	// "lms-cosmos/simapp/simd/cmd"
	"github.com/Leave-Management-System/lms-cosmos/simapp/lmsd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
