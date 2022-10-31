package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdRun)
}

var cmdRun = &cobra.Command{
	Use:   "run [exercise]",
	Short: "Run a single exercise",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := exercises.Run(args[0])
		if err != nil {
			color.Red("Compilation of %s failed! Compiler error message:\n\n%s", result.Exercise.Path, result.Err)
			color.Yellow("If you feel stuck, ask a hint by executing `golings hint %s`", result.Exercise.Name)
			os.Exit(1)
		} else {
			color.Green("Congratulations!\n\n")
			color.Green("Remove the 'I AM NOT DONE' from the file to keep going\n")
			color.Green("Here is the output of your program:\n\n")
			color.Cyan(result.Out)
			os.Exit(0)
		}
	},
}
