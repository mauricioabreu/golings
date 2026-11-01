package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// These tests build the real binary and drive it like a learner who installed
// it with `go install` and has no clone of the repository.

func buildBinary(t *testing.T) string {
	t.Helper()
	if testing.Short() {
		t.Skip("builds the binary and runs go; skipped in -short mode")
	}
	bin := filepath.Join(t.TempDir(), "golings")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
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
	expect(t, out, code, 0, "dev")

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
