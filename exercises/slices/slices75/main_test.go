// slices75
// Make the tests pass!

// I AM NOT DONE
//
// has checks whether a string is in the slice.
// Practices a linear search.
package main_test

import "testing"

func has(s []string, x string) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return len(s) > 0
}

func TestHas(t *testing.T) {
	s := []string{"go", "c"}
	if !has(s, "c") || has(s, "js") || has(nil, "go") {
		t.Errorf("has works incorrectly")
	}
}
