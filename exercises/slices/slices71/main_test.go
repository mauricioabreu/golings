// slices71
// Make the tests pass!

// I AM NOT DONE
//
// isEmpty must return true for a nil slice as well as for an empty one.
// len(nil) == 0.
package main_test

import "testing"

func isEmpty(s []int) bool {
	return s == nil
}

func TestIsEmpty(t *testing.T) {
	if !isEmpty(nil) || !isEmpty([]int{}) || isEmpty([]int{1}) {
		t.Errorf("isEmpty works incorrectly")
	}
}
