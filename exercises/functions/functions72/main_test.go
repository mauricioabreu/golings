// functions72
// Make the tests pass!

// I AM NOT DONE
//
// logLine must add a prefix and pass all the arguments on to format.
// Right now the arguments are lost.
// Practices building a new argument slice and passing it on with ...
package main_test

import (
	"strings"
	"testing"
)

func format(parts ...string) string {
	return strings.Join(parts, " ")
}

func logLine(prefix string, args ...string) string {
	return format(prefix)
}

func TestLogLine(t *testing.T) {
	if got := logLine("[info]", "server", "started"); got != "[info] server started" {
		t.Errorf("logLine = %q", got)
	}
	if got := logLine("[warn]"); got != "[warn]" {
		t.Errorf("logLine without args = %q", got)
	}
}
