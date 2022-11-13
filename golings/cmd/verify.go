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
		exs, err := exercises.List()
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		bar := progressbar.NewOptions(
			len(exs),
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
		bar.RenderBlank()

		for _, e := range exs {
			bar.Describe(fmt.Sprintf("Running %s", e.Name))
			result, _ := exercises.Run(e.Name)
			bar.Add(1)
			if result.Err != "" {
				fmt.Print("\n\n")
				color.Cyan("Failed to compile the exercise %s\n\n", e.Path)
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
