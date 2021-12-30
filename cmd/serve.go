package cmd

import (
	"github.com/kcraley/habiliti/internal/server"
	"github.com/kcraley/habiliti/pkg/terraform"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	address string
	port    string
)

// newServeCommand returns a new command which serves the application
func newServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "starts and serves the Habiliti server",
		Long:  `starts and serves the Habiliti server`,
		RunE:  serveCmdFuncE,
	}

	cmd.Flags().StringVarP(&address, "address", "a", "0.0.0.0", "the IP address that the application is being served")
	cmd.Flags().StringVarP(&port, "port", "p", "7000", "the port that the application is being served on")

	return cmd
}

func serveCmdFuncE(cmd *cobra.Command, args []string) error {
	log.Infof("starting application server on %s:%s...", address, port)

	// Create and serve the application server
	tfReg := terraform.NewRegistry(&terraform.RegistryOptions{})
	app := server.New(&server.Options{
		Address:           address,
		Port:              port,
		TerraformRegistry: tfReg,
	})
	if err := app.ListenAndServe(); err != nil {
		log.Fatalf("an error occurred serving the application: %s", err)
		return err
	}

	return nil
}
