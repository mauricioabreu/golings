package exercises

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

// DefaultTimeout bounds how long an exercise may run once it has been built.
// It is generous: it only has to catch infinite loops and goroutines that wait
// forever, not slow machines.
const DefaultTimeout = 30 * time.Second

// ErrTimeout means the built exercise was still running when time ran out.
var ErrTimeout = errors.New("the exercise didn't finish in time")

type Result struct {
	Exercise Exercise
	Out      string
	Err      string
	// BuildFailed is set when the exercise did not compile, as opposed to
	// compiling and then failing (or crashing) while it ran.
	BuildFailed bool
	// TimedOut is set when the exercise was stopped after Timeout.
	TimedOut bool
	Timeout  time.Duration
}

// Run builds and runs the exercise, giving up on it after DefaultTimeout.
func (e Exercise) Run() (Result, error) {
	return e.RunContext(context.Background(), DefaultTimeout)
}

// RunContext builds the exercise and runs the result. Building is cancelled
// together with ctx; running is also stopped after timeout.
//
// The exercise is built first and the program executed directly, instead of
// using `go run` or `go test`: those start the program as a child of the go
// command, and stopping the go command would leave the program running.
func (e Exercise) RunContext(ctx context.Context, timeout time.Duration) (Result, error) {
	result := Result{Exercise: e}

	if local, err := LocalGoVersion(); err == nil && !e.Supports(local) {
		return result, fmt.Errorf("%w: %s needs Go %s, found %s", ErrGoVersionTooOld, e.Name, e.GoVersion, local)
	}

	dir, err := os.MkdirTemp("", "golings-*")
	if err != nil {
		return result, err
	}
	defer os.RemoveAll(dir) // nolint

	binary := filepath.Join(dir, "exercise")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	build := exec.CommandContext(ctx, "go", buildArgs(e, binary)...)
	var buildOut, buildErr bytes.Buffer
	build.Stdout, build.Stderr = &buildOut, &buildErr
	if err := build.Run(); err != nil {
		result.Out, result.Err, result.BuildFailed = buildOut.String(), buildErr.String(), true
		if ctx.Err() != nil {
			return result, ctx.Err()
		}
		return result, err
	}

	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	run := exec.CommandContext(runCtx, binary, runArgs(e)...)
	if e.Mode == "test" {
		// `go test` runs the test binary in the package folder.
		run.Dir = filepath.Dir(e.Path)
	}
	// Don't wait for output pipes held open by anything the program started.
	run.WaitDelay = 2 * time.Second
	var stdout, stderr bytes.Buffer
	run.Stdout, run.Stderr = &stdout, &stderr

	err = run.Run()
	result.Out, result.Err = stdout.String(), stderr.String()

	switch {
	case err != nil && ctx.Err() != nil:
		return result, ctx.Err()
	case err != nil && runCtx.Err() != nil:
		result.TimedOut, result.Timeout = true, timeout
		return result, fmt.Errorf("%w (stopped after %s)", ErrTimeout, timeout)
	}
	return result, err
}

// buildArgs compiles the exercise into the file out: a program for compile
// mode, a test binary (with the race detector, as before) for test mode.
func buildArgs(e Exercise, out string) []string {
	path := fmt.Sprintf("./%s", e.Path)
	if e.Mode == "compile" {
		return []string{"build", "-o", out, path}
	}
	return []string{"test", "-c", "-race", "-o", out, path}
}

func runArgs(e Exercise) []string {
	if e.Mode == "compile" {
		return nil
	}
	return []string{"-test.v"}
}
