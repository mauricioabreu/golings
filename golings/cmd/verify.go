package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdVerify)
}

var cmdVerify = &cobra.Command{
	Use:   "verify",
	Short: "Verify all exercises",
	Run: func(cmd *cobra.Command, args []string) {
		verified, err := exercises.Verify()
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		bar := progressbar.NewOptions(
			3,
			progressbar.OptionSetWidth(50),
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionSetPredictTime(false),
			progressbar.OptionSetElapsedTime(false),
			progressbar.OptionSetDescription("[cyan][reset] Progress:"),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "[yellow]=[reset]",
				SaucerHead:    "[yellow]>[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		)

		for v := range verified {
			bar.Describe(fmt.Sprintf("Running %s", v.Exercise.Name))
			bar.Add(1)
			if v.Err != "" {
				fmt.Print("\n\n")
				color.Cyan("Failed to compile the exercise %s\n", v.Exercise.Path)
				color.Red("Check the error: %s", v.Err)
				break
			}
		}
	},
}
