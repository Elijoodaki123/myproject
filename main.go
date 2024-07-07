package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	app "github.com/your_username/my-cosmos-app/app"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myappd",
		Short: "My Cosmos SDK Application",
	}

	rootCmd.AddCommand(
		server.StartCmd(func(logger log.Logger, db dbm.DB) server.Application {
			return app.NewMyApp(logger, db)
		}, app.DefaultNodeHome),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
