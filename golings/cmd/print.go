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

	progress, done, total, err := exercises.Progress(infoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		color.Blue("Progress: %d/%d (%.2f%%)\n\n", done, total, progress*100)
	}

	if done == total {
		color.Green("Congratulations!!\n")
		color.Green("You have completed all %d of the currently available exercises.", total)
		color.Blue("If you enjoyed working through this introduction to Golang,")
		color.Blue("please give the github repository a star")
		color.White("> https://github.com/mauricioabreu/golings <\n\n\n")

		color.Yellow("To quit out of watch, please type `exit` and hit enter:")

		return;
	}

	exercise, err := exercises.NextPending(infoFile)
	if err != nil {
		color.Red("Failed to find next exercises")
	}

	result, err := exercise.Run()
	if err != nil {
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
			color.Red("exercise is still pending")
		}
	}
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
