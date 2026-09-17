// switch5
// Make the tests pass!

// I AM NOT DONE
//
// escape replaces a newline with \n, a tab with \t and a backslash with \\.
// Practices a switch on a rune inside a loop with a strings.Builder.
package main_test

import (
	"strings"
	"testing"
)

func escape(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case '\n':
			b.WriteString(`\n`)
		case '\t':
			b.WriteString(`\n`)
		}
	}
	return b.String()
}

func TestEscape(t *testing.T) {
	in := "a\tb\nc\\d"
	want := `a\tb\nc\\d`
	if got := escape(in); got != want {
		t.Errorf("escape(%q) = %s, want %s", in, got, want)
	}
}
