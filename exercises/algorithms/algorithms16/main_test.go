// algorithms16
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: greedy counting. Tasks are denoted by letters; identical tasks
// must be executed with a gap of at least n ticks. Return the minimum
// number of ticks (including idle time).
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func leastInterval(tasks []byte, n int) int {
	return 0
}

func TestLeastInterval(t *testing.T) {
	cases := []struct {
		tasks   []byte
		n, want int
	}{
		{[]byte("AAABBB"), 2, 8},
		{[]byte("AAABBB"), 0, 6},
		{[]byte("AAAAAABCDEFG"), 2, 16},
		{[]byte("A"), 5, 1},
		{nil, 3, 0},
	}
	for _, c := range cases {
		if got := leastInterval(c.tasks, c.n); got != c.want {
			t.Errorf("leastInterval(%q, %d) = %d, want %d", c.tasks, c.n, got, c.want)
		}
	}
}
