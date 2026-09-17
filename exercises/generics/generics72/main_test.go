// generics72
// Make the tests pass!

// I AM NOT DONE
//
// Equal compares two slices element by element.
// Practices comparable in a generic comparison.
package main_test

import "testing"

func Equal[T comparable](a, b []T) bool {
	if len(a) > len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestEqual(t *testing.T) {
	if !Equal([]int{1, 2}, []int{1, 2}) || Equal([]string{"a"}, []string{"a", "b"}) {
		t.Errorf("Equal works incorrectly")
	}
}
