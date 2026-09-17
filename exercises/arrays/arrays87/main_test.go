// arrays87
// Make the tests pass!

// I AM NOT DONE
//
// makeCounters must return three independent counters.
// Every pointer points at the same variable.
// Copying pointers does not copy the data.
package main_test

import "testing"

func makeCounters() [3]*int {
	var out [3]*int
	n := 0
	for i := range out {
		out[i] = &n
	}
	return out
}

func TestMakeCounters(t *testing.T) {
	c := makeCounters()
	*c[0] += 1
	*c[1] += 10
	if *c[0] != 1 || *c[1] != 10 || *c[2] != 0 {
		t.Errorf("counters = %d %d %d", *c[0], *c[1], *c[2])
	}
}
