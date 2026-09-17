// if7
// Make the tests pass!

// I AM NOT DONE
//
// halfDone must return true once at least half of the tasks are done.
// Right now the condition holds even with nothing done.
// Untyped integer constants divide as integers.
package main_test

import "testing"

func halfDone(done, total int) bool {
	if float64(done)/float64(total) >= 1/2 {
		return true
	}
	return false
}

func TestHalfDone(t *testing.T) {
	cases := []struct {
		done, total int
		want        bool
	}{{0, 10, false}, {4, 10, false}, {5, 10, true}, {3, 5, true}, {1, 3, false}}
	for _, c := range cases {
		if got := halfDone(c.done, c.total); got != c.want {
			t.Errorf("halfDone(%d, %d) = %v, want %v", c.done, c.total, got, c.want)
		}
	}
}
