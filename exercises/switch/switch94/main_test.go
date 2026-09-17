// switch94
// Make the tests pass!

// I AM NOT DONE
//
// compareWords compares two numbers and returns "less", "equal" or "greater".
// Practices a tagless switch with three outcomes.
package main_test

import "testing"

func compareWords(a, b int) string {
	switch {
	case a <= b:
		return "less"
	case a >= b:
		return "greater"
	}
	return "equal"
}

func TestCompareWords(t *testing.T) {
	cases := [][2]int{{1, 2}, {2, 2}, {3, 2}}
	want := []string{"less", "equal", "greater"}
	for i, c := range cases {
		if got := compareWords(c[0], c[1]); got != want[i] {
			t.Errorf("compareWords(%d, %d) = %s, want %s", c[0], c[1], got, want[i])
		}
	}
}
