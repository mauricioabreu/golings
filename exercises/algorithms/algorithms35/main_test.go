// algorithms35
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: binary search on a predicate. Versions 1..n; starting from
// some version, all versions are bad. Find the first bad one, calling
// isBad as rarely as possible. If there are no bad ones — return n+1.
// Expected asymptotics: O(log n) calls, O(1) space.
package main_test

import "testing"

func firstBad(n int, isBad func(int) bool) int {
	return 0
}

func TestFirstBad(t *testing.T) {
	cases := []struct{ n, bad int }{{5, 4}, {1, 1}, {10, 11}, {2147483647, 2147483646}, {100, 1}}
	for _, c := range cases {
		calls := 0
		got := firstBad(c.n, func(v int) bool { calls++; return v >= c.bad })
		if got != c.bad {
			t.Errorf("firstBad(n=%d) = %d, want %d", c.n, got, c.bad)
		}
		if calls > 40 {
			t.Errorf("firstBad(n=%d) made %d calls", c.n, calls)
		}
	}
}
