// arrays48
// Make the tests pass!

// I AM NOT DONE
//
// allOn returns true only when every switch is on.
// Practices checking every element of an array.
package main_test

import "testing"

func allOn(sw [4]bool) bool {
	for _, on := range sw {
		if on {
			return false
		}
	}
	return true
}

func TestAllOn(t *testing.T) {
	if !allOn([4]bool{true, true, true, true}) || allOn([4]bool{true, false, true, true}) {
		t.Errorf("allOn works incorrectly")
	}
}
