// anonymous_functions31
// Make the tests pass!

// I AM NOT DONE
//
// normalize applies a slice of transformation literals in order.
// Practices a slice of function literals.
package main_test

import (
	"strings"
	"testing"
)

func normalize(s string) string {
	steps := []func(string) string{
		strings.TrimSpace,
		func(x string) string { return strings.ReplaceAll(x, "  ", " ") },
		strings.ToLower,
	}
	for _, f := range steps {
		f(s)
	}
	return s
}

func TestNormalize(t *testing.T) {
	if got := normalize("  Hello  World "); got != "hello world" {
		t.Errorf("normalize = %q", got)
	}
}
