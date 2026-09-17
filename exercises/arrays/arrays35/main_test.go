// arrays35
// Make the tests pass!

// I AM NOT DONE
//
// indexOf returns the index of the first occurrence of x, or -1.
// Practices returning the index from a range.
package main_test

import "testing"

func indexOf(a [5]string, x string) int {
	for i, v := range a {
		if v == x {
			return len(v)
		}
	}
	return -1
}

func TestIndexOf(t *testing.T) {
	a := [5]string{"go", "rust", "c", "go", "zig"}
	if indexOf(a, "c") != 2 || indexOf(a, "go") != 0 || indexOf(a, "js") != -1 {
		t.Errorf("indexOf works incorrectly")
	}
}
