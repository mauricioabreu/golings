// algorithms89
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: one-dimensional DP. A digit string encodes letters:
// "1" — A, ..., "26" — Z. Count the number of ways to decode it.
// A leading zero makes decoding impossible.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func numDecodings(s string) int {
	return 0
}

func TestNumDecodings(t *testing.T) {
	cases := map[string]int{"12": 2, "226": 3, "06": 0, "0": 0, "": 0, "10": 1, "27": 1, "1111": 5, "100": 0}
	for in, want := range cases {
		if got := numDecodings(in); got != want {
			t.Errorf("numDecodings(%q) = %d, want %d", in, got, want)
		}
	}
}
