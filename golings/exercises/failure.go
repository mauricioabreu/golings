package exercises

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type FailureKind int

const (
	// CompileFailure means the Go compiler rejected the exercise.
	CompileFailure FailureKind = iota + 1
	// TestFailure means the exercise compiled but its tests failed.
	TestFailure
	// RuntimeFailure means the exercise compiled but crashed or exited non-zero.
	RuntimeFailure
	// TimeoutFailure means the exercise was stopped because it ran for too long.
	TimeoutFailure
)

// String describes the failure in a few words, for lists: "does not compile".
func (k FailureKind) String() string {
	switch k {
	case CompileFailure:
		return "does not compile"
	case TestFailure:
		return "has failing tests"
	case TimeoutFailure:
		return "does not finish"
	default:
		return "crashes"
	}
}

// Diagnostic is a single compiler message such as `main.go:10:6: undefined: x`.
type Diagnostic struct {
	File    string
	Line    int
	Col     int // 0 when the compiler gave no column
	Message string
}

// Failure is the learner-facing interpretation of a failed Result.
type Failure struct {
	Kind        FailureKind
	Diagnostics []Diagnostic  // only for CompileFailure
	Output      string        // everything else worth showing, with toolchain noise removed
	Timeout     time.Duration // for TimeoutFailure: how long the exercise was given
}

// Compiler messages start at the beginning of a line; indented file:line
// pairs are t.Log output from tests.
var diagnosticRegex = regexp.MustCompile(`^(?:vet: )?(\S.*?\.go):(\d+):(?:(\d+):)? (.+)$`)

// Failure interprets the output of a failed run.
func (r Result) Failure() Failure {
	combined := r.Err + "\n" + r.Out

	if r.TimedOut {
		return Failure{Kind: TimeoutFailure, Output: cleanOutput(combined), Timeout: r.Timeout}
	}

	if r.BuildFailed {
		if diagnostics, rest := parseDiagnostics(combined); len(diagnostics) > 0 {
			return Failure{Kind: CompileFailure, Diagnostics: diagnostics, Output: rest}
		}
		return Failure{Kind: CompileFailure, Output: cleanOutput(combined)}
	}

	// What the program printed comes first, then the panic, as it happened.
	ran := condensePanics(cleanOutput(r.Out + "\n" + r.Err))
	if r.Exercise.Mode == "test" {
		return Failure{Kind: TestFailure, Output: ran}
	}
	return Failure{Kind: RuntimeFailure, Output: ran}
}

func parseDiagnostics(output string) ([]Diagnostic, string) {
	var diagnostics []Diagnostic
	var rest []string
	for line := range strings.SplitSeq(output, "\n") {
		m := diagnosticRegex.FindStringSubmatch(line)
		if m == nil {
			rest = append(rest, line)
			continue
		}
		d := Diagnostic{File: m[1], Message: m[4]}
		d.Line, _ = strconv.Atoi(m[2])
		d.Col, _ = strconv.Atoi(m[3])
		diagnostics = append(diagnostics, d)
	}
	return diagnostics, cleanOutput(strings.Join(rest, "\n"))
}

// cleanOutput drops toolchain lines that only add noise for a learner and
// trims blank lines around what is left.
func cleanOutput(output string) string {
	var kept []string
	for line := range strings.SplitSeq(output, "\n") {
		switch {
		case strings.HasPrefix(line, "# "), // package header: `# command-line-arguments`
			line == "FAIL",
			strings.HasPrefix(line, "FAIL\t"),
			strings.HasPrefix(line, "ok  \t"):
			continue
		}
		kept = append(kept, line)
	}
	return strings.TrimSpace(strings.Join(kept, "\n"))
}

// frameLine matches the second line of a stack frame: a tab, a file and a line.
var frameLine = regexp.MustCompile(`^\t(.+\.go):(\d+)(?: \+0x[0-9a-f]+)?$`)

// condensePanics shortens what Go prints for a panic or a fatal error. The
// message stays, and so do the stack frames in the learner's code (shown as
// `at function (file:line)`, with the file relative to the working directory);
// frames inside the runtime and the testing package, the goroutine numbers and
// the memory addresses are left out.
func condensePanics(output string) string {
	cwd, _ := os.Getwd()
	lines := strings.Split(output, "\n")

	var out []string
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		switch {
		case strings.HasPrefix(line, "[signal "):
			continue
		case strings.HasPrefix(line, "panic: "), strings.HasPrefix(line, "fatal error: "):
			line = strings.TrimSuffix(strings.TrimSuffix(line, " [recovered, repanicked]"), " [recovered]")
		case isGoroutineHeader(line):
			// The frames follow as pairs: the function, then an indented file:line.
			for i+2 < len(lines) && lines[i+1] != "" {
				function, location := lines[i+1], lines[i+2]
				i += 2
				m := frameLine.FindStringSubmatch(location)
				if m == nil || isToolchainFrame(function, m[1]) {
					continue
				}
				out = append(out, fmt.Sprintf("  at %s (%s:%s)", shortFunction(function), relativeTo(cwd, m[1]), m[2]))
			}
			continue
		}
		out = append(out, line)
	}
	return strings.TrimSpace(strings.Join(out, "\n"))
}

// isGoroutineHeader matches lines such as `goroutine 8 [running]:`.
func isGoroutineHeader(line string) bool {
	return strings.HasPrefix(line, "goroutine ") && strings.HasSuffix(line, ":") && strings.Contains(line, " [")
}

func isToolchainFrame(function, file string) bool {
	file = filepath.ToSlash(file)
	return strings.HasPrefix(function, "created by ") ||
		strings.HasPrefix(function, "panic(") ||
		strings.HasPrefix(function, "testing.") ||
		strings.HasPrefix(function, "runtime.") ||
		strings.Contains(file, "/src/runtime/") ||
		strings.Contains(file, "/src/testing/")
}

// shortFunction turns `command-line-arguments_test.greeting(0xc00001)` into `greeting`.
func shortFunction(function string) string {
	if open := strings.Index(function, "("); open >= 0 {
		function = function[:open]
	}
	if dot := strings.LastIndex(function, "."); dot >= 0 {
		function = function[dot+1:]
	}
	return function
}

func relativeTo(dir, file string) string {
	if dir == "" {
		return file
	}
	if rel, err := filepath.Rel(dir, file); err == nil && !strings.HasPrefix(rel, "..") {
		return filepath.ToSlash(rel)
	}
	return file
}
