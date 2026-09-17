// arrays83
// Make the tests pass!

// I AM NOT DONE
//
// rotateLeft shifts the elements one place left, and the first one goes to the end.
// Practices copying array elements.
package main_test

import "testing"

func rotateLeft(a [5]int) [5]int {
	var out [5]int
	for i := range a {
		out[i] = a[(i+2)%len(a)]
	}
	return out
}

func TestRotateLeft(t *testing.T) {
	if got := rotateLeft([5]int{1, 2, 3, 4, 5}); got != [5]int{2, 3, 4, 5, 1} {
		t.Errorf("rotateLeft = %v", got)
	}
}
