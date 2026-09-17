// functions36
// Make the tests pass!

// I AM NOT DONE
//
// factorial must compute n! recursively (0! = 1).
// Right now the result is always 0.
// Practices the base case of a recursion.
package main_test

import "testing"

func factorial(n int) int {
	if n == 0 {
		return 0
	}
	return n * factorial(n-1)
}

func TestFactorial(t *testing.T) {
	cases := map[int]int{0: 1, 1: 1, 5: 120, 10: 3628800}
	for in, want := range cases {
		if got := factorial(in); got != want {
			t.Errorf("factorial(%d) = %d, want %d", in, got, want)
		}
	}
}
