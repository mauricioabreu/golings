// if50
// Make the tests pass!

// I AM NOT DONE
//
// tax: up to 10000 is 0%, up to 50000 is 10%, above that 20% of the whole amount.
// Right now the rates are applied to the wrong ranges.
// Practices an else if chain with rising thresholds.
package main_test

import "testing"

func tax(income int) int {
	if income < 10000 {
		return 0
	} else if income < 50000 {
		return income * 20 / 100
	}
	return income * 10 / 100
}

func TestTax(t *testing.T) {
	cases := map[int]int{5000: 0, 10000: 0, 20000: 2000, 50000: 5000, 100000: 20000}
	for in, want := range cases {
		if got := tax(in); got != want {
			t.Errorf("tax(%d) = %d, want %d", in, got, want)
		}
	}
}
