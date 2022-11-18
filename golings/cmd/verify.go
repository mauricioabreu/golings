package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func VerifyCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "verify",
		Short: "Verify all exercises",
		Run: func(cmd *cobra.Command, args []string) {
			allExercises, err := exercises.List(infoFile)
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
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
				os.Exit(1)
			}

			for _, exercise := range allExercises {
				bar.Describe(fmt.Sprintf("Running %s", exercise.Name))
				result, _ := exercise.Run()
				bar.Add(1) // nolint

				if result.Err != "" {
					fmt.Print("\n\n")
					color.Cyan("Failed to compile the exercise %s\n\n", exercise.Path)
					color.White("Check the output below: \n\n")
					color.Red(result.Err)
					color.Red(result.Out)
					os.Exit(1)
				}
			}

			color.Green("Congratulations!!!")
			color.Green("You passed all the exercises")
		},
	}
}
