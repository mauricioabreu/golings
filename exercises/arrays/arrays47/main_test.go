// arrays47
// Make the tests pass!

// I AM NOT DONE
//
// prev returns the index of the previous position in a ring array of 5 elements.
// For position 0 it comes out as -1 and the program panics.
// The remainder of a negative number when indexing an array.
package main_test

import "testing"

var ring = [5]string{"a", "b", "c", "d", "e"}

func prev(i int) string {
	return ring[(i-1)%len(ring)]
}

func TestPrev(t *testing.T) {
	if prev(3) != "c" || prev(0) != "e" || prev(1) != "a" {
		t.Errorf("prev works incorrectly")
	}
}
