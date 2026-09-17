// variables30
// Make the tests pass!

// I AM NOT DONE
//
// divider must return a string of n '-' characters.
// The code does not compile: a string cannot be multiplied by a number.
// Practices string values through the strings package.
package main_test

import (
	"strings"
	"testing"
)

func divider(n int) string {
	return "-" * n
}

func TestDivider(t *testing.T) {
	if got := divider(5); got != "-----" {
		t.Errorf("divider(5) = %q, want %q", got, "-----")
	}
	if got := divider(0); got != "" {
		t.Errorf("divider(0) = %q, want empty", got)
	}
}
