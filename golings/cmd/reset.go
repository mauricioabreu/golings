package cmd

import (
	"bufio"
	"errors"
	"io/fs"
	"strings"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/workspace"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ResetCmd restores exercises below root from src.
func ResetCmd(src fs.FS, root string) *cobra.Command {
	var yes bool

	resetCmd := &cobra.Command{
		Use:   "reset <exercise name>",
		Short: "Restore an exercise to its original state",
		Long: `Overwrites the exercise file with the original version that comes with golings,
discarding your changes. The "I AM NOT DONE" marker comes back with it.`,
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			if !yes {
				color.Yellow("This overwrites your changes to %s with the original exercise. Continue? [y/N]", name)
				answer, _ := bufio.NewReader(cmd.InOrStdin()).ReadString('\n')
				if answer = strings.ToLower(strings.TrimSpace(answer)); answer != "y" && answer != "yes" {
					color.White("Nothing was changed.")
					return errors.New("reset cancelled")
				}
			}

			path, err := workspace.Reset(src, root, name)
			if errors.Is(err, exercises.ErrExerciseNotFound) {
				color.White("No exercise found for '%s'", name)
				return err
			}
			if err != nil {
				color.Red("Failed to reset %s: %s", name, err)
				return err
			}

			color.Green("Reset %s to its original state.", path)
			return nil
		},
	}
	resetCmd.Flags().BoolVarP(&yes, "yes", "y", false, "do not ask for confirmation")

	return resetCmd
}
