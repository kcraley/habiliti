package cmd

import (
	"github.com/kcraley/habiliti/internal/server"
	"github.com/kcraley/habiliti/pkg/terraform"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// newServeCommand returns a new command which serves the application
func newServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "starts and serves the Habiliti server",
		Long:  `starts and serves the Habiliti server`,
		RunE:  serveCmdFuncE,
	}

	cmd.Flags().StringVarP(&config.Address, "address", "a", "0.0.0.0", "the IP address that the application is being served")
	cmd.Flags().BoolVarP(&config.EnableLogin, "enable-login", "l", true, "controls enabling the login endpoint")
	cmd.Flags().BoolVarP(&config.EnableModules, "enable-modules", "m", true, "controls enabling the modules endpoint")
	cmd.Flags().BoolVarP(&config.EnableProviders, "enable-providers", "r", true, "controls enabling the providers endpoint")
	cmd.Flags().StringVarP(&config.Endpoint, "endpoint", "e", "/v1", "the root endpoint to serve the Terraform registry")
	cmd.Flags().StringVarP(&config.Port, "port", "p", "7000", "the port that the application is being served on")

	return cmd
}

func serveCmdFuncE(cmd *cobra.Command, args []string) error {
	log.Infof("starting application server on %s:%s...", config.Address, config.Port)

	// Create and serve the application server
	tfReg := terraform.NewRegistry(&terraform.RegistryOptions{
		EnableLogin:     config.EnableLogin,
		EnableModules:   config.EnableModules,
		EnableProviders: config.EnableProviders,
	})
	app := server.New(&server.Options{
		Address:           config.Address,
		Endpoint:          config.Endpoint,
		Port:              config.Port,
		TerraformRegistry: tfReg,
	})
	if err := app.ListenAndServe(config.Address, config.Port); err != nil {
		log.Fatalf("an error occurred serving the application: %s", err)
		return err
	}

	return nil
}
