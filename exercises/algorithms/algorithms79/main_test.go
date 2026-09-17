// algorithms79
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: working with the digits of a number. Check whether an integer
// reads the same left to right and right to left, without converting it
// to a string.
// Expected asymptotics: O(log n) time, O(1) space.
package main_test

import "testing"

func isPalindromeNumber(x int) bool {
	return false
}

func TestIsPalindromeNumber(t *testing.T) {
	cases := map[int]bool{121: true, -121: false, 10: false, 0: true, 1221: true, 12321: true, 1000021: false}
	for in, want := range cases {
		if got := isPalindromeNumber(in); got != want {
			t.Errorf("isPalindromeNumber(%d) = %v, want %v", in, got, want)
		}
	}
}
