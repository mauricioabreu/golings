// anonymous_functions46
// Make the tests pass!

// I AM NOT DONE
//
// toggler returns a closure flipping the state on every call.
// Practices a closure over a bool.
package main_test

import "testing"

func toggler() func() bool {
	on := false
	return func() bool {
		on = true
		return on
	}
}

func TestToggler(t *testing.T) {
	tg := toggler()
	if !tg() || tg() || !tg() {
		t.Errorf("toggler works incorrectly")
	}
}
