package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func RunCmd(infoFile string) *cobra.Command {
	return &cobra.Command{
		Use:   "run next | <exercise name>",
		Short: "Run a single exercise",
		Long: `example next pending exercise : golings run next
example specific exercise : golings run variables1`,
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

			spinner := RunSpinner(exercise.Name)

			if errors.Is(err, exercises.ErrExerciseNotFound) {
				color.White("No exercise found for '%s'", args[0])
				return err
			}

			result, err := exercise.Run()

			spinner.Close()

			if err != nil {
				color.Cyan("Failed to compile the exercise %s\n\n", result.Exercise.Path)
				color.White("Check the output below: \n\n")
				color.Red(result.Err)
				color.Red(result.Out)
				color.Yellow("If you feel stuck, ask a hint by executing `golings hint %s`", result.Exercise.Name)
			} else {
				color.Green("✅ Successfully tested %s!\n\n", result.Exercise.Path)
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

func RunSpinner(exercise string) *progressbar.ProgressBar {
	spinner := progressbar.NewOptions(
		-1, // a negative number makes turns the progress bar into a spinner
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription(color.WhiteString("Running exercise: %s", exercise)),
		progressbar.OptionOnCompletion(func() {
			color.White("\nRunning complete!\n\n")
		}),
	)
	go func() {
		for x := 0; x < 100; x++ {
			spinner.Add(1) // nolint
			time.Sleep(250 * time.Millisecond)
		}
	}()

	return spinner
}
