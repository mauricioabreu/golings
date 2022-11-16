package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "golings",
		Short:   "Learn go through interactive exercises",
		Version: version,
	}

	rootCmd.AddCommand(cmdHint)
	rootCmd.AddCommand(ListCmd("info.toml"))
	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(cmdVerify)

	return rootCmd
}

func Execute(version string) {
	NewRootCmd(version).Execute()
}
