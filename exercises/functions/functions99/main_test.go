// functions99
// Make the tests pass!

// I AM NOT DONE
//
// render builds text by passing a strings.Builder to a helper function.
// The code panics: "illegal use of non-zero Builder copied by value".
// Some types must not be copied once they are in use.
package main_test

import (
	"strings"
	"testing"
)

func writeLine(b strings.Builder, s string) {
	b.WriteString(s)
	b.WriteString("\n")
}

func render(lines []string) string {
	var b strings.Builder
	b.WriteString("BEGIN\n")
	for _, l := range lines {
		writeLine(b, l)
	}
	b.WriteString("END")
	return b.String()
}

func TestRender(t *testing.T) {
	want := "BEGIN\nx\ny\nEND"
	if got := render([]string{"x", "y"}); got != want {
		t.Errorf("render = %q, want %q", got, want)
	}
}
