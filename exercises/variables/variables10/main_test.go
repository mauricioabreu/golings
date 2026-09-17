// variables10
// Make the tests pass!

// I AM NOT DONE
//
// toggle must invert a boolean value.
// Practices logical negation and the zero value of a bool.
package main_test

import "testing"

func toggle(on bool) bool {
	return on
}

func TestToggle(t *testing.T) {
	var light bool
	light = toggle(light)
	if !light {
		t.Errorf("toggle(false) = false, want true")
	}
	if toggle(light) {
		t.Errorf("toggle(true) = true, want false")
	}
}
