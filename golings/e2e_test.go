package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// These tests build the real binary and drive it like a learner who installed
// it with `go install` and has no clone of the repository. The binary is stamped
// with a version, the way GoReleaser does, because a plain `go build` inside a
// git checkout reports a version derived from the commit.

func buildBinary(t *testing.T) string {
	t.Helper()
	if testing.Short() {
		t.Skip("builds the binary and runs go; skipped in -short mode")
	}
	bin := filepath.Join(t.TempDir(), "golings")
	if out, err := exec.Command("go", "build", "-ldflags", "-X main.version=e2e-test", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return bin
}

// run executes the binary in dir and returns its combined output and exit code.
func run(t *testing.T, bin, dir string, args ...string) (string, int) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "NO_COLOR=1")
	out, err := cmd.CombinedOutput()
	var exitErr *exec.ExitError
	switch {
	case err == nil:
		return string(out), 0
	case errors.As(err, &exitErr):
		return string(out), exitErr.ExitCode()
	default:
		t.Fatalf("running %v: %v", args, err)
		return "", -1
	}
}

func expect(t *testing.T, out string, code, wantCode int, wantInOutput ...string) {
	t.Helper()
	if code != wantCode {
		t.Errorf("exit code = %d, want %d\n%s", code, wantCode, out)
	}
	for _, want := range wantInOutput {
		if !strings.Contains(out, want) {
			t.Errorf("output does not contain %q\n%s", want, out)
		}
	}
}

func TestLearnerJourney(t *testing.T) {
	bin := buildBinary(t)
	home := t.TempDir()

	out, code := run(t, bin, home, "--version")
	expect(t, out, code, 0, "e2e-test")

	out, code = run(t, bin, home, "init")
	expect(t, out, code, 0, "Created golings-exercises", "cd golings-exercises")
	ws := filepath.Join(home, "golings-exercises")

	out, code = run(t, bin, home, "init")
	expect(t, out, code, 1, "already exists")

	out, code = run(t, bin, ws, "init")
	expect(t, out, code, 1, "already a golings workspace")

	out, code = run(t, bin, ws, "doctor")
	expect(t, out, code, 0, "Exercise files: all present")

	out, code = run(t, bin, home, "list")
	expect(t, out, code, 1) // not in a workspace: info.toml is missing

	out, code = run(t, bin, ws, "run", "variables1")
	expect(t, out, code, 1, "Failed to compile the exercise", "syntax error", "^")

	out, code = run(t, bin, ws, "run", "if1")
	expect(t, out, code, 1, "Tests failed for the exercise")

	out, code = run(t, bin, ws, "hint", "variables1")
	expect(t, out, code, 0, "Variables must have names")

	out, code = run(t, bin, ws, "hint", "nope")
	expect(t, out, code, 1, "exercise not found")

	out, code = run(t, bin, ws, "verify")
	expect(t, out, code, 1, "Failed to compile the exercise")

	// The solution stays hidden until the exercise is done, or on request.
	out, code = run(t, bin, ws, "solution", "variables1")
	expect(t, out, code, 1, "not done yet")
	if strings.Contains(out, "var x = 5") {
		t.Errorf("the solution leaked:\n%s", out)
	}
	out, code = run(t, bin, ws, "solution", "variables1", "--reveal")
	expect(t, out, code, 0, "var x = 5", `"x has the value %d"`)

	// Fix variables1 the way a learner would.
	file := filepath.Join(ws, "exercises", "variables", "variables1", "main.go")
	original, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	fixed := strings.ReplaceAll(string(original), "var = 5", "var x = 5")
	fixed = strings.ReplaceAll(fixed, "// I AM NOT DONE\n", "")
	if err := os.WriteFile(file, []byte(fixed), 0o644); err != nil {
		t.Fatal(err)
	}
	out, code = run(t, bin, ws, "run", "variables1")
	expect(t, out, code, 0, "Successfully tested")
	out, code = run(t, bin, ws, "solution", "variables1")
	expect(t, out, code, 0, "var x = 5")

	// init must never hand out the solutions.
	if _, err := os.Stat(filepath.Join(ws, "solutions")); err == nil {
		t.Errorf("init wrote a solutions folder into the learner's workspace")
	}

	// Reset brings back the broken original, marker included.
	out, code = run(t, bin, ws, "reset", "variables1", "--yes")
	expect(t, out, code, 0, "Reset exercises/variables/variables1/main.go")
	restored, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if string(restored) != string(original) {
		t.Errorf("reset did not restore the original exercise:\n%s", restored)
	}
	out, code = run(t, bin, ws, "run", "variables1")
	expect(t, out, code, 1, "Failed to compile the exercise")

	out, code = run(t, bin, ws, "reset", "nope", "--yes")
	expect(t, out, code, 1, "No exercise found for 'nope'")
}

// lockedBuffer collects the output of a process that is still running.
type lockedBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *lockedBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *lockedBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

func waitFor(t *testing.T, out *lockedBuffer, want string) {
	t.Helper()
	deadline := time.Now().Add(60 * time.Second)
	for !strings.Contains(out.String(), want) {
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for %q in:\n%s", want, out.String())
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func TestWatchReactsToSavedFiles(t *testing.T) {
	bin := buildBinary(t)
	home := t.TempDir()
	if out, code := run(t, bin, home, "init"); code != 0 {
		t.Fatalf("init failed:\n%s", out)
	}
	ws := filepath.Join(home, "golings-exercises")

	cmd := exec.Command(bin, "watch")
	cmd.Dir = ws
	cmd.Env = append(os.Environ(), "NO_COLOR=1")
	out := &lockedBuffer{}
	cmd.Stdout, cmd.Stderr = out, out
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = cmd.Process.Kill() })

	// The first exercise is broken...
	waitFor(t, out, "Failed to compile the exercise")

	// ...and saving a fix, the way an editor does (write a temp file, rename it
	// over the original), makes watch run it again.
	file := filepath.Join(ws, "exercises", "variables", "variables1", "main.go")
	original, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	fixed := strings.ReplaceAll(string(original), "var = 5", "var x = 5")
	temp := file + ".tmp"
	if err := os.WriteFile(temp, []byte(fixed), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(temp, file); err != nil {
		t.Fatal(err)
	}
	waitFor(t, out, "Here is the output of your program")

	if _, err := stdin.Write([]byte("exit\n")); err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case err := <-done:
		if err != nil {
			t.Errorf("watch exited with an error after `exit`: %v\n%s", err, out.String())
		}
	case <-time.After(10 * time.Second):
		t.Errorf("watch did not exit after `exit`:\n%s", out.String())
	}
}
