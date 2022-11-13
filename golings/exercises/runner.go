package exercises

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Result struct {
	Exercise Exercise
	Out      string
	Err      string
}

func Run(name string) (Result, error) {
	exercise, err := Find(name)
	if err != nil {
		return Result{}, err
	}

	args := BuildArgs(exercise)
	cmd := exec.Command("go", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	return Result{Exercise: exercise, Out: stdout.String(), Err: stderr.String()}, err
}

func BuildArgs(e Exercise) []string {
	args := []string{}
	if e.Mode == "compile" {
		args = append(args, "run")
	} else {
		args = append(args, "test", "-v")
	}

	args = append(args, fmt.Sprintf("./%s", e.Path))
	return args
}
