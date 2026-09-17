// variables78
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the largest int32 value.
// Practices the predefined constants of the math package.
package main_test

import (
	"math"
	"testing"
)

func maxInt32() int32 {
	var limit int32 = math.MaxInt16
	return limit
}

func TestMaxInt32(t *testing.T) {
	if got := maxInt32(); got != 2147483647 {
		t.Errorf("maxInt32() = %d, want 2147483647", got)
	}
}
