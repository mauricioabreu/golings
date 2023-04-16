package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/spf13/cobra"
)

func HintCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "hint <exercise name>",
		Short: "Get a hint for an exercise",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			var exercise exercises.Exercise
			var err error
			if args[0] == "next" {
				exercise, err = exercises.NextPending(infoFile)
			} else {
				exercise, err = exercises.Find(args[0], infoFile)
			}

			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}
			color.Yellow(exercise.Hint)
		},
	}
}
