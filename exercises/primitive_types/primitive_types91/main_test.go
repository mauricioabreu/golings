// primitive_types91
// Make the tests pass!

// I AM NOT DONE
//
// unsignedShift must shift the bits of an int32 right, filling the vacated
// top bits with zeros, the way >>> does in Java.
// For signed types >> is an arithmetic shift.
package main_test

import "testing"

func unsignedShift(x int32, n uint) int32 {
	return x >> n
}

func TestUnsignedShift(t *testing.T) {
	if got := unsignedShift(16, 2); got != 4 {
		t.Errorf("unsignedShift(16, 2) = %d", got)
	}
	if got := unsignedShift(-1, 28); got != 15 {
		t.Errorf("unsignedShift(-1, 28) = %d, want 15", got)
	}
	if got := unsignedShift(-8, 1); got != 2147483644 {
		t.Errorf("unsignedShift(-8, 1) = %d, want 2147483644", got)
	}
}
