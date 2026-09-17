// anonymous_functions39
// Make the tests pass!

// I AM NOT DONE
//
// track wraps a function and counts its calls in a captured variable.
// Practices changing an outer variable from a literal.
package main_test

import "testing"

func track(f func(), calls *int) func() {
	return func() {
		*calls = 1
		f()
	}
}

func TestTrack(t *testing.T) {
	n := 0
	g := track(func() {}, &n)
	g()
	g()
	g()
	if n != 3 {
		t.Errorf("calls = %d, want 3", n)
	}
}
