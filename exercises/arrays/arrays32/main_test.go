// arrays32
// Make the tests pass!

// I AM NOT DONE
//
// swapEnds swaps the first and the last element.
// Practices a multiple assignment of array elements.
package main_test

import "testing"

func swapEnds(a [4]int) [4]int {
	a[0] = a[3]
	return a
}

func TestSwapEnds(t *testing.T) {
	if got := swapEnds([4]int{1, 2, 3, 4}); got != [4]int{4, 2, 3, 1} {
		t.Errorf("swapEnds = %v", got)
	}
}
