// arrays30
// Make the tests pass!

// I AM NOT DONE
//
// capacity returns the capacity of an array, which always equals its length.
// Practices cap on arrays.
package main_test

import "testing"

func capacity(a [8]byte) int {
	return cap(a) * 2
}

func TestCapacity(t *testing.T) {
	if got := capacity([8]byte{}); got != 8 {
		t.Errorf("capacity = %d, want 8", got)
	}
}
