// functions45
// Make the tests pass!

// I AM NOT DONE
//
// isEven and isOdd are defined in terms of each other for non-negative n.
// The base cases and the calls are muddled up.
// Practices mutually recursive functions.
package main_test

import "testing"

func isEven(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

func TestParity(t *testing.T) {
	for n := 0; n < 10; n++ {
		if isEven(n) != (n%2 == 0) || isOdd(n) != (n%2 == 1) {
			t.Errorf("parity of %d: isEven=%v isOdd=%v", n, isEven(n), isOdd(n))
		}
	}
}
