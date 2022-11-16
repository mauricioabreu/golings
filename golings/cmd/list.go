package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/mauricioabreu/golings/golings/ui"
	"github.com/spf13/cobra"
)

func ListCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all exercises",
		Run: func(cmd *cobra.Command, args []string) {
			exs, err := exercises.List(infoFile)
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}
			ui.PrintList(os.Stdout, exs)
		},
	}
}
