package cmd

import (
	"os"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/ui"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func ListCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:           "list",
		Short:         "List all exercises",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			exs, err := exercises.List(infoFile)
			if err != nil {
				color.Red(err.Error())
				return err
			}
			ui.PrintList(os.Stdout, exs)
			return nil
		},
	}
}
