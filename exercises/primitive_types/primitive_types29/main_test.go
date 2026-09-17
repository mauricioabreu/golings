// primitive_types29
// Make the tests pass!

// I AM NOT DONE
//
// dashes must replace every space with a dash.
// Practices strings.Replace and strings.ReplaceAll.
package main_test

import (
	"strings"
	"testing"
)

func dashes(s string) string {
	return strings.Replace(s, " ", "-", 1)
}

func TestDashes(t *testing.T) {
	if got := dashes("a b c d"); got != "a-b-c-d" {
		t.Errorf("dashes = %q, want a-b-c-d", got)
	}
}
