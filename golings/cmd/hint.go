package cmd

import (
	"github.com/cleonte/golings/golings/exercises"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func HintCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:           "hint <exercise name>",
		Short:         "Get a hint for an exercise",
		Args:          cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var exercise exercises.Exercise
			var err error
			if args[0] == "next" {
				exercise, err = exercises.NextPending(infoFile)
			} else {
				exercise, err = exercises.Find(args[0], infoFile)
			}

			if err != nil {
				color.Red(err.Error())
				return err
			}
			color.Yellow(exercise.Hint)
			return nil
		},
	}
}
