// range78
// Make the tests pass!

// I AM NOT DONE
//
// countLikeFirst counts how many elements, the first one aside, equal the first.
// Practices a range over a subslice.
package main_test

import "testing"

func countLikeFirst(s []string) int {
	if len(s) == 0 {
		return 0
	}
	n := 0
	for _, v := range s {
		if v == s[0] {
			n++
		}
	}
	return n
}

func TestCountLikeFirst(t *testing.T) {
	if got := countLikeFirst([]string{"a", "b", "a", "a"}); got != 2 {
		t.Errorf("countLikeFirst = %d, want 2", got)
	}
}
