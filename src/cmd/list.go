package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/src/exercises"
	"github.com/mauricioabreu/golings/src/printer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdList)
}

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List all exercises",
	Run: func(cmd *cobra.Command, args []string) {
		exs, err := exercises.List()
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		printer.PrintList(os.Stdout, exs)
	},
}
