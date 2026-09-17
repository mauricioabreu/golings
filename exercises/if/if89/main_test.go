// if89
// Make the tests pass!

// I AM NOT DONE
//
// compareVersion compares versions (major, minor) and returns -1, 0 or 1.
// Practices a nested comparison over several fields.
package main_test

import "testing"

func compareVersion(aMaj, aMin, bMaj, bMin int) int {
	if aMin < bMin {
		return -1
	}
	if aMin > bMin {
		return 1
	}
	return 0
}

func TestCompareVersion(t *testing.T) {
	cases := []struct{ a1, a2, b1, b2, want int }{
		{1, 2, 1, 3, -1}, {2, 0, 1, 9, 1}, {1, 5, 1, 5, 0}, {1, 9, 2, 0, -1},
	}
	for _, c := range cases {
		if got := compareVersion(c.a1, c.a2, c.b1, c.b2); got != c.want {
			t.Errorf("compare(%d.%d, %d.%d) = %d, want %d", c.a1, c.a2, c.b1, c.b2, got, c.want)
		}
	}
}
