// arrays75
// Make the tests pass!

// I AM NOT DONE
//
// indexOfMax must return the index of the largest element.
// Practices the two variables of a range over an array.
package main_test

import "testing"

func indexOfMax(a [5]int) int {
	best := 0
	for v, i := range a {
		if v > a[best] {
			best = i
		}
	}
	return best
}

func TestIndexOfMax(t *testing.T) {
	if got := indexOfMax([5]int{3, 9, 2, 9, 0}); got != 1 {
		t.Errorf("indexOfMax = %d, want 1", got)
	}
}
