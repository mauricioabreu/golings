// maps71
// Make the tests pass!

// I AM NOT DONE
//
// firstUnique returns the index of the first non-repeating byte of a string, or -1.
// Practices two passes: counting and searching.
package main_test

import "testing"

func firstUnique(s string) int {
	counts := map[byte]int{}
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}
	for c, n := range counts {
		if n == 1 {
			return int(c)
		}
	}
	return -1
}

func TestFirstUnique(t *testing.T) {
	cases := map[string]int{"leetcode": 0, "loveleetcode": 2, "aabb": -1, "": -1}
	for in, want := range cases {
		if got := firstUnique(in); got != want {
			t.Errorf("firstUnique(%q) = %d, want %d", in, got, want)
		}
	}
}
