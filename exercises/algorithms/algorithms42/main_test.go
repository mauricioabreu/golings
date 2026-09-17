// algorithms42
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: base-26 numeral system. Convert a column title
// ("A", "B", ..., "Z", "AA", ...) to its number.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func titleToNumber(s string) int {
	return 0
}

func TestTitleToNumber(t *testing.T) {
	cases := map[string]int{"A": 1, "AB": 28, "ZY": 701, "": 0, "FXSHRXW": 2147483647}
	for in, want := range cases {
		if got := titleToNumber(in); got != want {
			t.Errorf("titleToNumber(%q) = %d, want %d", in, got, want)
		}
	}
}
