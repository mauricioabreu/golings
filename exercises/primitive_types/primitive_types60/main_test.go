// primitive_types60
// Make the tests pass!

// I AM NOT DONE
//
// steps counts how many 0.1 steps it takes to go from 0 to 1.
// The condition != 1.0 never holds exactly, and the loop hits its safety limit.
// A loop must not be driven by an exact float comparison.
package main_test

import "testing"

func steps() int {
	n := 0
	for x := 0.0; x != 1.0 && n < 1000; x += 0.1 {
		n++
	}
	return n
}

func TestSteps(t *testing.T) {
	if got := steps(); got != 10 {
		t.Errorf("steps() = %d, want 10", got)
	}
}
