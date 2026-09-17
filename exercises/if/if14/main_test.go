// if14
// Make the tests pass!

// I AM NOT DONE
//
// label must return "big" for numbers over 100 and "small" otherwise.
// The result is always "small".
// := inside an if block creates a new variable.
package main_test

import "testing"

func label(n int) string {
	result := "small"
	if n > 100 {
		result := "big"
		_ = result
	}
	return result
}

func TestLabel(t *testing.T) {
	if got := label(500); got != "big" {
		t.Errorf("label(500) = %q, want big", got)
	}
	if got := label(5); got != "small" {
		t.Errorf("label(5) = %q, want small", got)
	}
}
