package exercises

import (
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

	if r.Exercise.Mode == "test" {
		return Failure{Kind: TestFailure, Output: cleanOutput(combined)}
	}
	return Failure{Kind: RuntimeFailure, Output: cleanOutput(combined)}
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
