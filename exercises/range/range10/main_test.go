// range10
// Make the tests pass!

// I AM NOT DONE
//
// countTrue counts the flags that are on.
// Practices a range over a map[string]bool.
package main_test

import "testing"

func countTrue(flags map[string]bool) int {
	n := 0
	for _, on := range flags {
		if !on {
			n++
		}
	}
	return n
}

func TestCountTrue(t *testing.T) {
	if got := countTrue(map[string]bool{"a": true, "b": false, "c": true}); got != 2 {
		t.Errorf("countTrue = %d, want 2", got)
	}
}
