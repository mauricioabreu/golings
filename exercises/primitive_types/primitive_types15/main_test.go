// primitive_types15
// Make the tests pass!

// I AM NOT DONE
//
// fieldCount must return the number of fields in a comma separated line.
// Practices strings.Split.
package main_test

import (
	"strings"
	"testing"
)

func fieldCount(line string) int {
	return len(strings.Split(line, ";"))
}

func TestFieldCount(t *testing.T) {
	if got := fieldCount("a,b,c"); got != 3 {
		t.Errorf("fieldCount(a,b,c) = %d, want 3", got)
	}
	if got := fieldCount("x"); got != 1 {
		t.Errorf("fieldCount(x) = %d, want 1", got)
	}
}
