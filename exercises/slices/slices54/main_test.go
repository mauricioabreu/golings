// slices54
// Make the tests pass!

// I AM NOT DONE
//
// total adds up every element of a slice.
// Practices a range over a slice.
package main_test

import "testing"

func total(s []int) int {
	sum := 0
	for _, v := range s {
		sum = v
	}
	return sum
}

func TestTotal(t *testing.T) {
	if got := total([]int{5, 10, 15}); got != 30 {
		t.Errorf("total = %d, want 30", got)
	}
}
