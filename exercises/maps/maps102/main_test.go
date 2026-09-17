// maps102
// Make the tests pass!

// I AM NOT DONE
//
// sameFrequency checks that every character of a string occurs the same number of times.
// Practices frequencies of frequencies, a map of the values.
package main_test

import "testing"

func sameFrequency(s string) bool {
	counts := map[rune]int{}
	for _, r := range s {
		counts[r]++
	}
	freqs := map[int]bool{}
	for _, n := range counts {
		freqs[n] = n > 1
	}
	return len(freqs) > 0
}

func TestSameFrequency(t *testing.T) {
	cases := map[string]bool{"aabb": true, "abc": true, "aab": false, "": true, "xxyyzzz": false}
	for in, want := range cases {
		if got := sameFrequency(in); got != want {
			t.Errorf("sameFrequency(%q) = %v, want %v", in, got, want)
		}
	}
}
