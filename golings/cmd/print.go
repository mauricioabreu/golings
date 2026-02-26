package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/mauricioabreu/golings/golings/exercises"
	"github.com/mauricioabreu/golings/golings/ui"
)

func PrintHint(infoFile string) {
	ClearScreen()
	exercise, err := exercises.NextPending(infoFile)
	if err != nil {
		color.Red("Failed to find next exercises")
	}
	color.Yellow(exercise.Hint)
}

func PrintList(infoFile string) {
	ClearScreen()
	exs, err := exercises.List(infoFile)
	if err != nil {
		color.Red("Failed to list exercises")
	}
	ui.PrintList(os.Stdout, exs)
}

func RunNextExercise(infoFile string) {
	ClearScreen()

	allExs, err := exercises.List(infoFile)
	if err != nil {
		color.Red("Failed to list exercises: %v", err)
		return
	}

	var target exercises.Exercise
	var lastResult exercises.Result
	found := false

	// Iterate through all exercises to find the first one that isn't actually "done"
	for _, ex := range allExs {
		res, _ := ex.Run()
		// An exercise is only truly done if it compiles/passes AND the comment is gone
		if res.Err != "" || ex.State() == exercises.Pending {
			target = ex
			lastResult = res
			found = true
			break
		}
	}

	// Update progress based on actual completion
	progress, done, total, err := exercises.Progress(infoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		color.Blue("Progress: %d/%d (%.2f%%)\n\n", done, total, progress*100)
	}

	if !found {
		color.Green("Congratulations!!\n")
		color.Green("You have completed all %d of the currently available exercises.", total)
		color.Blue("If you enjoyed working through this introduction to Golang,")
		color.Blue("please give the github repository a star")
		color.White("> https://github.com <\n\n\n")
		color.Yellow("To quit out of watch, please type `exit` and hit enter:")
		return
	}

	// If there's a compilation/test error, show it immediately
	if lastResult.Err != "" {
		color.Cyan("Failed to compile the exercise %s\n\n", lastResult.Exercise.Path)
		color.White("Check the output below: \n\n")
		color.Red(lastResult.Err)
		color.Red(lastResult.Out)
		color.Yellow("If you feel stuck, ask a hint by executing `golings hint %s`", target.Name)
		return
	}

	// If it passes but the comment is still there
	color.Green("Successfully tested %s!\n\n", target.Path)
	color.Green("Congratulations!\n\n")
	color.Green("Here is the output of your program:\n\n")
	color.Cyan(lastResult.Out)
	color.White("Remove the 'I AM NOT DONE' from the file to keep going\n")
	color.Red("exercise is still pending")
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			color.Red("Clear terminal command error")
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			color.Red("Clear terminal command error")
		}
	}
}
