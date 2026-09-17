// algorithms80
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a stack. Evaluate an expression consisting of integers, '+', '-',
// parentheses, and spaces. A unary minus is allowed before a number or
// a parenthesis.
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import "testing"

func calculate(s string) int {
	return 0
}

func TestCalculate(t *testing.T) {
	cases := map[string]int{
		"1 + 1":               2,
		" 2-1 + 2 ":           3,
		"(1+(4+5+2)-3)+(6+8)": 23,
		"-(2+3)":              -5,
		"42":                  42,
		"- (3 - (4 - 10))":    -9,
		"2147483647 + 1":      2147483648,
	}
	for in, want := range cases {
		if got := calculate(in); got != want {
			t.Errorf("calculate(%q) = %d, want %d", in, got, want)
		}
	}
}
