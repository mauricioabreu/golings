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

func (e Exercise) Run() (Result, error) {
	args := BuildArgs(e)
	cmd := exec.Command("go", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	return Result{Exercise: e, Out: stdout.String(), Err: stderr.String()}, cmd.Run()
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
