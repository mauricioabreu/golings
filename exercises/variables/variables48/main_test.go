// variables48
// Make the tests pass!

// I AM NOT DONE
//
// wrapIndex must map any index into the range [0, n).
// For negative indexes the result falls outside the range.
// Practices the sign of the % operator in Go.
package main_test

import "testing"

func wrapIndex(i, n int) int {
	return i % n
}

func TestWrapIndex(t *testing.T) {
	cases := []struct{ i, n, want int }{
		{0, 5, 0}, {7, 5, 2}, {-1, 5, 4}, {-5, 5, 0}, {-12, 5, 3},
	}
	for _, c := range cases {
		if got := wrapIndex(c.i, c.n); got != c.want {
			t.Errorf("wrapIndex(%d, %d) = %d, want %d", c.i, c.n, got, c.want)
		}
	}
}
