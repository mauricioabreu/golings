package cmd

import (
	"errors"
	"fmt"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func VerifyCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:           "verify",
		Short:         "Verify all exercises",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			allExercises, err := exercises.List(infoFile)
			if err != nil {
				color.Red(err.Error())
				return err
			}

			bar := progressbar.NewOptions(
				len(allExercises),
				progressbar.OptionSetWidth(50),
				progressbar.OptionEnableColorCodes(true),
				progressbar.OptionSetPredictTime(false),
				progressbar.OptionSetElapsedTime(false),
				progressbar.OptionSetDescription("[cyan][reset] Running exercises"),
				progressbar.OptionSetTheme(progressbar.Theme{
					Saucer:        "[yellow]=[reset]",
					SaucerHead:    "[yellow]>[reset]",
					SaucerPadding: " ",
					BarStart:      "[",
					BarEnd:        "]",
				}),
			)
			if err := bar.RenderBlank(); err != nil {
				color.Red(err.Error())
				return err
			}

			for _, exercise := range allExercises {
				bar.Describe(fmt.Sprintf("Running %s", exercise.Name))
				result, err := exercise.Run()
				bar.Add(1) // nolint

				if errors.Is(err, exercises.ErrGoVersionTooOld) {
					fmt.Print("\n")
					color.Yellow("Skipped: %s", err)
					continue
				}

				if result.Err != "" {
					fmt.Print("\n\n")
					color.Cyan("Failed to compile the exercise %s\n\n", exercise.Path)
					color.White("Check the output below: \n\n")
					color.Red(result.Err)
					color.Red(result.Out)
					return errors.New("verification failed")
				}
			}

			color.Green("Congratulations!!!")
			color.Green("You passed all the exercises")
			return nil
		},
	}
}
