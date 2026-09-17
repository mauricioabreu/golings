// arrays20
// Make the tests pass!

// I AM NOT DONE
//
// fullName joins the parts of a name with spaces.
// Practices turning an array into a slice for strings.Join.
package main_test

import (
	"strings"
	"testing"
)

func fullName(parts [3]string) string {
	return strings.Join(parts[:], "")
}

func TestFullName(t *testing.T) {
	if got := fullName([3]string{"Anna", "Maria", "Ivanova"}); got != "Anna Maria Ivanova" {
		t.Errorf("fullName = %q", got)
	}
}
