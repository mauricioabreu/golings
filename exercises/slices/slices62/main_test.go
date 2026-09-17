// slices62
// Make the tests pass!

// I AM NOT DONE
//
// csv joins the fields with commas.
// Practices strings.Join on a []string.
package main_test

import (
	"strings"
	"testing"
)

func csv(fields []string) string {
	return strings.Join(fields, ", ")
}

func TestCSV(t *testing.T) {
	if got := csv([]string{"a", "b", "c"}); got != "a,b,c" {
		t.Errorf("csv = %q", got)
	}
}
