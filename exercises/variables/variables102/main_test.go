// variables102
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the share of finished tasks as a percentage (0..100).
// For 1 out of 3 the expected result is 33, but it comes out as 0.
// Practices the order of operations in integer division.
package main_test

import "testing"

func percentDone(done, total int) int {
	if total == 0 {
		return 0
	}
	return done / total * 100
}

func TestPercentDone(t *testing.T) {
	cases := []struct{ done, total, want int }{
		{1, 3, 33}, {2, 3, 66}, {3, 3, 100}, {0, 5, 0}, {1, 0, 0},
	}
	for _, c := range cases {
		if got := percentDone(c.done, c.total); got != c.want {
			t.Errorf("percentDone(%d, %d) = %d, want %d", c.done, c.total, got, c.want)
		}
	}
}
