package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdHint)
}

var cmdHint = &cobra.Command{
	Use:   "hint",
	Short: "Get a hint for an exercise",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		exercise, err := exercises.Find(args[0], "info.toml")
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		color.Yellow(exercise.Hint)
	},
}
