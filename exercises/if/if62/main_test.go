// if62
// Make the tests pass!

// I AM NOT DONE
//
// isEmpty must return true for a nil slice as well as for an empty one.
// Practices checking the length instead of comparing with nil.
package main_test

import "testing"

func isEmpty(items []string) bool {
	if items == nil {
		return true
	}
	return false
}

func TestIsEmpty(t *testing.T) {
	if !isEmpty(nil) || !isEmpty([]string{}) {
		t.Errorf("nil and empty slices should be empty")
	}
	if isEmpty([]string{"x"}) {
		t.Errorf("[x] is not empty")
	}
}
