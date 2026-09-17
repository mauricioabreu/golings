// variables71
// Make the tests pass!

// I AM NOT DONE
//
// This function must return Big*4/8, where Big = 2^62.
// The intermediate result does not fit in an int64.
// Untyped constants are evaluated with arbitrary precision.
package main_test

import "testing"

const Big = 1 << 62

func halfOfDouble() int {
	x := Big
	return x * 4 / 8
}

func TestHalfOfDouble(t *testing.T) {
	if got := halfOfDouble(); got != 1<<61 {
		t.Errorf("halfOfDouble() = %d, want %d", got, 1<<61)
	}
}
