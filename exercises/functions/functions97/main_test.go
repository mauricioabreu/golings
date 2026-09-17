// functions97
// Make the tests pass!

// I AM NOT DONE
//
// sumTo(n, acc) must return 1+2+...+n, building the sum up in acc.
// Practices recursion with an accumulator.
package main_test

import "testing"

func sumTo(n, acc int) int {
	if n == 0 {
		return 0
	}
	return sumTo(n-1, acc)
}

func TestSumTo(t *testing.T) {
	cases := map[int]int{0: 0, 1: 1, 10: 55, 1000: 500500}
	for n, want := range cases {
		if got := sumTo(n, 0); got != want {
			t.Errorf("sumTo(%d, 0) = %d, want %d", n, got, want)
		}
	}
}
