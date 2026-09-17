// slices43
// Make the tests pass!

// I AM NOT DONE
//
// zeros must return a slice of n zeros.
// Practices make([]T, len).
package main_test

import "testing"

func zeros(n int) []int {
	return make([]int, 0)
}

func TestZeros(t *testing.T) {
	z := zeros(3)
	if len(z) != 3 || z[0] != 0 || z[2] != 0 {
		t.Errorf("zeros(3) = %v", z)
	}
}
