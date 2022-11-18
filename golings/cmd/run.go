package cmd

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/spf13/cobra"
)

func RunCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:           "run [exercise]",
		Short:         "Run a single exercise",
		Args:          cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := exercises.Run(args[0], infoFile)
			if errors.Is(err, exercises.ErrExerciseNotFound) {
				color.White("No exercise found for '%s'", args[0])
			} else if err != nil {
				color.Cyan("Failed to compile the exercise %s\n\n", result.Exercise.Path)
				color.White("Check the output below: \n\n")
				color.Red(result.Err)
				color.Red(result.Out)
				color.Yellow("If you feel stuck, ask a hint by executing `golings hint %s`", result.Exercise.Name)
			} else {
				color.Green("Congratulations!\n\n")
				color.Green("Here is the output of your program:\n\n")
				color.Cyan(result.Out)
				if result.Exercise.State() == exercises.Pending {
					color.White("Remove the 'I AM NOT DONE' from the file to keep going\n")
					return fmt.Errorf("exercise is still pending")
				}
			}

			return err
		},
	}
}
