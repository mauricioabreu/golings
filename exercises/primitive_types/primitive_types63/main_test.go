// primitive_types63
// Make the tests pass!

// I AM NOT DONE
//
// shout must turn a string into upper case.
// Practices the functions of the strings package.
package main_test

import (
	"strings"
	"testing"
)

func shout(s string) string {
	return strings.ToTitle(strings.ToLower(s)) + "!"
}

func TestShout(t *testing.T) {
	if got := shout("go fast"); got != "GO FAST" {
		t.Errorf("shout = %q, want GO FAST", got)
	}
}
