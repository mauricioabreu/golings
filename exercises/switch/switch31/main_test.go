// switch31
// Make the tests pass!

// I AM NOT DONE
//
// insert takes a coin and returns the new balance; the coins 1, 2 and 5 are accepted,
// a 10 only while the balance is below 10, and anything else is rejected without changing the balance.
// Practices a switch with a condition inside a branch.
package main_test

import "testing"

func insert(balance, coin int) int {
	switch coin {
	case 1, 2, 5:
		return balance + coin
	default:
		return balance + coin
	}
}

func TestInsert(t *testing.T) {
	cases := [][3]int{{0, 5, 5}, {3, 2, 5}, {0, 10, 10}, {12, 10, 12}, {4, 3, 4}, {0, 50, 0}}
	for _, c := range cases {
		if got := insert(c[0], c[1]); got != c[2] {
			t.Errorf("insert(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
