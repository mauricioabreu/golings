// if99
// Make the tests pass!

// I AM NOT DONE
//
// needsReview: an order does NOT need a review only when it is paid AND
// the amount is below 10000. Every other case needs one.
// Practices negating a compound condition correctly.
package main_test

import "testing"

func needsReview(paid bool, total int) bool {
	if !paid && total >= 10000 {
		return true
	}
	return false
}

func TestNeedsReview(t *testing.T) {
	cases := []struct {
		paid  bool
		total int
		want  bool
	}{{true, 500, false}, {false, 500, true}, {true, 20000, true}, {false, 20000, true}}
	for _, c := range cases {
		if got := needsReview(c.paid, c.total); got != c.want {
			t.Errorf("needsReview(%v, %d) = %v, want %v", c.paid, c.total, got, c.want)
		}
	}
}
