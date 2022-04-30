package cmd

import (
	"log"

	"github.com/CanDgrmc/gotask/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server ",
	Run:   run,
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func run(cmd *cobra.Command, args []string) {
	var config *api.Configuration
	viper.Unmarshal(&config)

	server, err := api.NewServer(config)

	if err != nil {
		log.Fatal(err)
	}
	server.Start()
}
