package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/src/exercises"
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
		out, err := exercises.Run(args[0])
		fmt.Println(out)
		if err != nil {
			color.Red("Your exercise is failing: %s", err)
			os.Exit(1)
		}
		color.Green("Congratulations!")
		color.Green("Remove the 'I AM NOT DONE' from the file to keep going")
		os.Exit(0)
	},
}
