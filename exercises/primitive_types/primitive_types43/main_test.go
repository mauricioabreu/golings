// primitive_types43
// Make the tests pass!

// I AM NOT DONE
//
// label must return the string "item-7" for the number 7.
// Practices formatting integers with fmt.Sprintf.
package main_test

import (
	"fmt"
	"testing"
)

func label(n int) string {
	return fmt.Sprintf("item-%c", n)
}

func TestLabel(t *testing.T) {
	if got := label(7); got != "item-7" {
		t.Errorf("label(7) = %q, want item-7", got)
	}
}
