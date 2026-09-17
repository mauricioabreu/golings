// if65
// Make the tests pass!

// I AM NOT DONE
//
// divisible must report whether a divides by b with no remainder.
// Practices the remainder operator in a condition.
package main_test

import "testing"

func divisible(a, b int) bool {
	if a/b == 0 {
		return true
	}
	return false
}

func TestDivisible(t *testing.T) {
	cases := []struct {
		a, b int
		want bool
	}{{10, 5, true}, {10, 3, false}, {0, 7, true}, {3, 9, false}}
	for _, c := range cases {
		if got := divisible(c.a, c.b); got != c.want {
			t.Errorf("divisible(%d, %d) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
