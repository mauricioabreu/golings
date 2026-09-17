// if60
// Make the tests pass!

// I AM NOT DONE
//
// canPay must return true when there is enough money for the purchase (balance >= price).
// Practices the order of the operands in a comparison.
package main_test

import "testing"

func canPay(balance, price int) bool {
	if price >= balance {
		return true
	}
	return false
}

func TestCanPay(t *testing.T) {
	cases := [][3]int{{100, 50, 1}, {50, 50, 1}, {49, 50, 0}, {0, 1, 0}}
	for _, c := range cases {
		if got := canPay(c[0], c[1]); got != (c[2] == 1) {
			t.Errorf("canPay(%d, %d) = %v", c[0], c[1], got)
		}
	}
}
