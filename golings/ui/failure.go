package ui

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/fatih/color"
)

const tabWidth = 4

// write prints to w, colored when c is not nil. Write errors are ignored: this
// is terminal output and there is nothing useful to do about them.
func write(w io.Writer, c *color.Color, format string, args ...any) {
	if c == nil {
		_, _ = fmt.Fprintf(w, format, args...)
		return
	}
	_, _ = c.Fprintf(w, format, args...)
}

// PrintFailure explains why an exercise failed. Compiler errors are shown with
// the offending source line and a caret under the column.
func PrintFailure(w io.Writer, ex exercises.Exercise, f exercises.Failure) {
	cyan, red, white := color.New(color.FgCyan), color.New(color.FgRed), color.New(color.FgWhite)

	switch f.Kind {
	case exercises.CompileFailure:
		write(w, cyan, "Failed to compile the exercise %s\n\n", ex.Path)
	case exercises.TestFailure:
		write(w, cyan, "Tests failed for the exercise %s\n\n", ex.Path)
	case exercises.TimeoutFailure:
		write(w, cyan, "The exercise %s didn't finish within %s and was stopped\n\n", ex.Path, f.Timeout)
		write(w, nil, "Look for an infinite loop, or a goroutine that waits forever (a channel nobody sends to, a lock that is never released).\n\n")
	default:
		write(w, cyan, "The exercise %s failed while running\n\n", ex.Path)
	}

	for _, d := range f.Diagnostics {
		location := fmt.Sprintf("%s:%d", d.File, d.Line)
		if d.Col > 0 {
			location += fmt.Sprintf(":%d", d.Col)
		}
		write(w, red, "%s: %s\n", location, d.Message)
		printSnippet(w, d)
		write(w, nil, "\n")
	}

	if f.Output != "" {
		if f.Kind != exercises.CompileFailure {
			write(w, white, "Output:\n")
			write(w, nil, "\n")
		}
		write(w, red, "%s\n", f.Output)
		write(w, nil, "\n")
	}
}

func printSnippet(w io.Writer, d exercises.Diagnostic) {
	data, err := os.ReadFile(d.File)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")
	if d.Line < 1 || d.Line > len(lines) {
		return
	}

	source := strings.TrimRight(lines[d.Line-1], "\r")
	gutter := fmt.Sprintf("%d", d.Line)
	pad := strings.Repeat(" ", len(gutter))

	write(w, nil, "\n  %s | %s\n", gutter, expandTabs(source))
	if d.Col > 0 {
		prefix := source
		if d.Col-1 < len(prefix) {
			prefix = prefix[:d.Col-1]
		}
		// Columns count bytes, the caret has to line up with what is displayed.
		offset := len([]rune(expandTabs(prefix)))
		write(w, nil, "  %s | %s^\n", pad, strings.Repeat(" ", offset))
	}
}

func expandTabs(s string) string {
	return strings.ReplaceAll(s, "\t", strings.Repeat(" ", tabWidth))
}
