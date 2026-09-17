// anonymous_functions92
// Make the tests pass!

// I AM NOT DONE
//
// limit is computed by a function literal called right away.
// Practices an IIFE, an immediately invoked function expression.
package main_test

import "testing"

func limit() int {
	l := func() int {
		base := 10
		return base + 10
	}()
	return l
}

func TestLimit(t *testing.T) {
	if limit() != 100 {
		t.Errorf("limit = %d, want 100", limit())
	}
}
