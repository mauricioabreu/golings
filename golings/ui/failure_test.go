package ui_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/cleonte/golings/golings/ui"
	"github.com/fatih/color"
	. "github.com/onsi/gomega"
)

func render(t *testing.T, ex exercises.Exercise, f exercises.Failure) string {
	t.Helper()
	prev := color.NoColor
	color.NoColor = true
	t.Cleanup(func() { color.NoColor = prev })

	var out bytes.Buffer
	ui.PrintFailure(&out, ex, f)
	return out.String()
}

func writeSource(t *testing.T, source string) string {
	t.Helper()
	file := filepath.Join(t.TempDir(), "main.go")
	if err := os.WriteFile(file, []byte(source), 0o644); err != nil {
		t.Fatal(err)
	}
	return file
}

func TestPrintCompileFailureShowsSourceAndCaret(t *testing.T) {
	g := NewWithT(t)
	file := writeSource(t, "package main\n\nfunc main() {\n\tvar = 5\n}\n")
	f := exercises.Failure{
		Kind:        exercises.CompileFailure,
		Diagnostics: []exercises.Diagnostic{{File: file, Line: 4, Col: 6, Message: "syntax error"}},
	}

	got := render(t, exercises.Exercise{Path: "a/main.go"}, f)

	g.Expect(got).To(ContainSubstring("Failed to compile the exercise a/main.go"))
	g.Expect(got).To(ContainSubstring(file + ":4:6: syntax error"))
	// the tab is shown as four spaces and the caret sits under the `=` (column 6)
	g.Expect(got).To(ContainSubstring("  4 |     var = 5\n"))
	g.Expect(got).To(ContainSubstring("    |         ^\n"))
}

func TestPrintCompileFailureWithoutColumnOrSource(t *testing.T) {
	g := NewWithT(t)
	f := exercises.Failure{
		Kind:        exercises.CompileFailure,
		Diagnostics: []exercises.Diagnostic{{File: "does/not/exist.go", Line: 3, Message: "oops"}},
	}

	got := render(t, exercises.Exercise{Path: "a/main.go"}, f)

	g.Expect(got).To(ContainSubstring("does/not/exist.go:3: oops"))
	g.Expect(got).NotTo(ContainSubstring("^"))
}

func TestPrintLineOutOfRange(t *testing.T) {
	g := NewWithT(t)
	file := writeSource(t, "package main\n")
	f := exercises.Failure{
		Kind:        exercises.CompileFailure,
		Diagnostics: []exercises.Diagnostic{{File: file, Line: 99, Col: 1, Message: "oops"}},
	}

	got := render(t, exercises.Exercise{Path: "a/main.go"}, f)

	g.Expect(got).To(ContainSubstring(":99:1: oops"))
	g.Expect(got).NotTo(ContainSubstring("|"))
}

func TestPrintTestAndRuntimeFailures(t *testing.T) {
	g := NewWithT(t)

	tests := render(t, exercises.Exercise{Path: "a/main_test.go"}, exercises.Failure{Kind: exercises.TestFailure, Output: "--- FAIL: TestX"})
	g.Expect(tests).To(ContainSubstring("Tests failed for the exercise a/main_test.go"))
	g.Expect(tests).To(ContainSubstring("--- FAIL: TestX"))
	g.Expect(tests).NotTo(ContainSubstring("Failed to compile"))

	runtime := render(t, exercises.Exercise{Path: "a/main.go"}, exercises.Failure{Kind: exercises.RuntimeFailure, Output: "panic: boom"})
	g.Expect(runtime).To(ContainSubstring("The exercise a/main.go failed while running"))
	g.Expect(runtime).To(ContainSubstring("panic: boom"))
}

func TestPrintTimeoutFailure(t *testing.T) {
	g := NewWithT(t)

	got := render(t, exercises.Exercise{Path: "a/main.go"}, exercises.Failure{Kind: exercises.TimeoutFailure, Timeout: 30 * time.Second, Output: "started"})

	g.Expect(got).To(ContainSubstring("The exercise a/main.go didn't finish within 30s and was stopped"))
	g.Expect(got).To(ContainSubstring("infinite loop"))
	g.Expect(got).To(ContainSubstring("started"))
}
