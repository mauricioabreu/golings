// arrays94
// Make the tests pass!

// I AM NOT DONE
//
// track remembers the state of an array before the changes and reports whether it changed.
// The snapshot is kept as a pointer and therefore changes along with the original.
// Assigning an array copies it, while taking its address does not.
package main_test

import "testing"

func track(state *[3]int, mutate func(*[3]int)) bool {
	before := state
	mutate(state)
	return *before != *state
}

func TestTrack(t *testing.T) {
	s := [3]int{1, 2, 3}
	if !track(&s, func(p *[3]int) { p[0] = 9 }) {
		t.Errorf("change not detected")
	}
	if track(&s, func(p *[3]int) {}) {
		t.Errorf("false change detected")
	}
}
