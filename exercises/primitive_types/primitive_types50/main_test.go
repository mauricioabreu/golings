// primitive_types50
// Make the tests pass!

// I AM NOT DONE
//
// area takes the sides as int32 values and must return the area as an int64.
// For large sides the result overflows before the conversion.
// Practices the order of the conversions in a multiplication.
package main_test

import "testing"

func area(w, h int32) int64 {
	return int64(w * h)
}

func TestArea(t *testing.T) {
	if got := area(3, 4); got != 12 {
		t.Errorf("area(3, 4) = %d", got)
	}
	if got := area(100000, 100000); got != 10000000000 {
		t.Errorf("area(1e5, 1e5) = %d, want 10000000000", got)
	}
}
