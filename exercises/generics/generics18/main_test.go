// generics18
// Make the tests pass!

// I AM NOT DONE
//
// Contains checks whether an element is in a slice.
// The code does not compile: == calls for another constraint.
// Practices the comparable constraint.
package main_test

import "testing"

func Contains[T any](s []T, x T) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func TestContains(t *testing.T) {
	if !Contains([]string{"a", "b"}, "b") || Contains([]int{1, 2}, 3) {
		t.Errorf("Contains works incorrectly")
	}
}
