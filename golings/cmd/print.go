package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/ui"
	"github.com/fatih/color"
)

func PrintHint(infoFile string) {
	ClearScreen()
	exercise, err := exercises.NextPending(infoFile)
	if err != nil {
		color.Red("Failed to find next exercises")
		return
	}
	color.Yellow(exercise.Hint)
}

func PrintList(infoFile string) {
	ClearScreen()
	exs, err := exercises.List(infoFile)
	if err != nil {
		color.Red("Failed to list exercises")
		return
	}
	ui.PrintList(os.Stdout, exs)
	printTopicProgress(infoFile)
}

// printTopicProgress prints the per-topic table below the exercise list.
func printTopicProgress(infoFile string) {
	topics, err := exercises.ProgressByTopic(infoFile)
	if err != nil || len(topics) == 0 {
		return
	}
	fmt.Println()
	ui.PrintTopicProgress(os.Stdout, topics)
}

// currentTopicLine describes how far along the learner is in a topic,
// e.g. "Topic: slices (2/5)". It is empty for exercises without a topic.
func currentTopicLine(infoFile string, exercise exercises.Exercise) string {
	if exercise.Topic() == "" {
		return ""
	}
	topics, err := exercises.ProgressByTopic(infoFile)
	if err != nil {
		return ""
	}
	for _, t := range topics {
		if t.Topic == exercise.Topic() {
			return fmt.Sprintf("Topic: %s (%d/%d)", t.Topic, t.Done, t.Total)
		}
	}
	return ""
}

func RunNextExercise(infoFile string) {
	ClearScreen()

	progress, done, total, err := exercises.Progress(infoFile)
	if err != nil {
		color.Red("Failed to read the exercises: %s", err)
		color.Yellow("Run golings from the folder that holds info.toml, create one with `golings init`, or try `golings doctor`.")
		return
	}
	color.Blue("Progress: %d/%d (%.2f%%)\n\n", done, total, progress*100)

	if skipped, err := exercises.Skipped(infoFile); err == nil && len(skipped) > 0 {
		color.Yellow("%d exercise(s) skipped: they need a newer Go than %s\n\n", len(skipped), localGo())
	}

	if total == 0 {
		color.Yellow("There are no exercises to run. Try `golings doctor` to check your setup.")
		return
	}

	if done == total {
		color.Green("Congratulations!!\n")
		color.Green("You have completed all %d of the currently available exercises.", total)
		color.Blue("If you enjoyed working through this introduction to Golang,")
		color.Blue("please give the github repository a star")
		color.White("> https://github.com/cleonte/golings <\n\n\n")

		color.Yellow("To quit out of watch, please type `exit` and hit enter:")

		return
	}

	exercise, err := exercises.NextPending(infoFile)
	if err != nil {
		color.Red("Failed to find next exercises")
		return
	}

	if line := currentTopicLine(infoFile, exercise); line != "" {
		color.Blue(line + "\n")
	}

	result, err := exercise.Run()
	if errors.Is(err, exercises.ErrGoVersionTooOld) {
		color.Yellow(err.Error())
		return
	}
	if err != nil {
		printFailure(result)
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

func localGo() string {
	v, err := exercises.LocalGoVersion()
	if err != nil {
		return "the installed version"
	}
	return v
}

// printFailure explains why the exercise failed and where to find help.
func printFailure(result exercises.Result) {
	ui.PrintFailure(color.Output, result.Exercise, result.Failure())
	color.Yellow("If you feel stuck, ask a hint by executing `golings hint %s`", result.Exercise.Name)
}
