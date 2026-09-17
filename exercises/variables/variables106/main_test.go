// variables106
// Make the tests pass!

// I AM NOT DONE
//
// tax must return 20% tax on a price in cents, truncated.
// The code does not compile: a fractional constant cannot take part in an int expression.
// Practices untyped constants and explicit conversions.
package main_test

import "testing"

const taxRate = 0.2

func tax(price int) int {
	return price * taxRate
}

func TestTax(t *testing.T) {
	cases := map[int]int{100: 20, 999: 199, 0: 0}
	for in, want := range cases {
		if got := tax(in); got != want {
			t.Errorf("tax(%d) = %d, want %d", in, got, want)
		}
	}
}
