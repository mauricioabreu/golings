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

	cmd := exec.Command("go", "run", fmt.Sprintf("./%s", exercise.Path))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	return Result{Exercise: exercise, Out: stdout.String(), Err: stderr.String()}, err
}
