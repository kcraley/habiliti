package cmd

import (
	"fmt"

	"github.com/kcraley/habiliti/internal/version"

	"github.com/spf13/cobra"
)

// newVersionCommand returns a new command which prints the application version
func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "prints the version and build information of the binary",
		Long:  `prints the version and build information of the binary`,
		Run:   versionCmdFunc,
	}

	return cmd
}

// versionCmd is the entrypoint for `hability version`
func versionCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println(version.String())
}
