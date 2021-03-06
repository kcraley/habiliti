package cmd

import (
	"os"

	"github.com/kcraley/habiliti/internal/configuration"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const appName = "habiliti"

var (
	// Application configuration
	config *configuration.Configuration = configuration.New()

	rootCmd = &cobra.Command{
		Use:   appName,
		Short: "a Terraform module and provider registry",
		Long: `a Terraform module and provider registry

The third-party Terraform module and provider registry.`,
	}
)

func initConfig() {
	// Marshal configuration from environment variables
	err := envconfig.Process(appName, config)
	if err != nil {
		log.Fatalf("failed parsing configuration: %v", err)
	}

	// Configure logging
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	if config.Verbose {
		log.SetLevel(log.DebugLevel)
	}
}

// Execute is the main entrypoint for the application
func Execute() {
	cobra.OnInitialize(initConfig)

	// Configure all persistent flags
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "enables addition verbose output for troubleshooting")

	// Add all subcommands to the root
	rootCmd.AddCommand(newVersionCommand())
	rootCmd.AddCommand(newServeCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
