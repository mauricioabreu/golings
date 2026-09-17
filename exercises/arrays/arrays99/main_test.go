// arrays99
// Make the tests pass!

// I AM NOT DONE
//
// squares returns an array of the squares of the indexes: [0 1 4 9 16].
// Practices writing to array elements in a loop.
package main_test

import "testing"

func squares() [5]int {
	var a [5]int
	for i := range a {
		a[i] = i * 2
	}
	return a
}

func TestSquares(t *testing.T) {
	if got := squares(); got != [5]int{0, 1, 4, 9, 16} {
		t.Errorf("squares = %v", got)
	}
}
