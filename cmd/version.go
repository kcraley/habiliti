package cmd

import (
	"fmt"

	"github.com/kcraley/habiliti/internal/version"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "prints the version and build information of the binary",
		Long:  `prints the version and build information of the binary`,
		Run:   versionCmdFunc,
	}
)

func init() {
	// Add `version` subcommand to `habiliti`
	rootCmd.AddCommand(versionCmd)
}

// versionCmd is the entrypoint for `hability version`
func versionCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println(version.String())
}
