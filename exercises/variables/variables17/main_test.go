// variables17
// Make the tests pass!

// I AM NOT DONE
//
// reset must zero the counter it is given.
// Right now it works on a copy of the value.
// Practices pointers and changing a variable through one.
package main_test

import "testing"

func reset(counter int) {
	counter = 0
}

func resetAndGet() int {
	hits := 42
	reset(hits)
	return hits
}

func TestReset(t *testing.T) {
	if got := resetAndGet(); got != 0 {
		t.Errorf("after reset hits = %d, want 0", got)
	}
}
