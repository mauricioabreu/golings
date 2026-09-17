// range41
// Make the tests pass!

// I AM NOT DONE
//
// countNewlines counts the newlines in a byte slice.
// Practices a range over a []byte.
package main_test

import "testing"

func countNewlines(data []byte) int {
	n := 0
	for _, b := range data {
		if b == 'n' {
			n++
		}
	}
	return n
}

func TestCountNewlines(t *testing.T) {
	if got := countNewlines([]byte("one\ntwo\nnine")); got != 2 {
		t.Errorf("countNewlines = %d, want 2", got)
	}
}
