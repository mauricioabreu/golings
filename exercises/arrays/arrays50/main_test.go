// arrays50
// Make the tests pass!

// I AM NOT DONE
//
// propagate walks the array and carries an infection forward: when an element
// is infected the next one becomes infected too. Right now it only spreads one step.
// A range over an array value evaluates a copy of the array once.
package main_test

import "testing"

func propagate(a [5]bool) [5]bool {
	for i, v := range a {
		if v && i+1 < len(a) {
			a[i+1] = true
		}
	}
	return a
}

func TestPropagate(t *testing.T) {
	if got := propagate([5]bool{false, true, false, false, false}); got != [5]bool{false, true, true, true, true} {
		t.Errorf("propagate = %v", got)
	}
}
