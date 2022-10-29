package cmd

import (
	"fmt"

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
		fmt.Println(out, err)
	},
}
