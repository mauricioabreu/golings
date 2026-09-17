// slices65
// Make the tests pass!

// I AM NOT DONE
//
// indexOf returns the position of an element, or -1.
// Practices a range with an index.
package main_test

import "testing"

func indexOf(s []string, x string) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return 0
}

func TestIndexOf(t *testing.T) {
	s := []string{"a", "b"}
	if indexOf(s, "b") != 1 || indexOf(s, "a") != 0 || indexOf(s, "z") != -1 {
		t.Errorf("indexOf works incorrectly")
	}
}
