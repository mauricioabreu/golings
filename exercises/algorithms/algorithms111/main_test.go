// algorithms111
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: bit manipulation. Check whether an integer is a power of two.
// Zero and negative numbers do not count as powers of two.
// Expected asymptotics: O(1) time, O(1) space.
package main_test

import "testing"

func isPowerOfTwo(x int) bool {
	return false
}

func TestIsPowerOfTwo(t *testing.T) {
	cases := map[int]bool{1: true, 16: true, 3: false, 0: false, -16: false, 1 << 62: true, 6: false}
	for in, want := range cases {
		if got := isPowerOfTwo(in); got != want {
			t.Errorf("isPowerOfTwo(%d) = %v, want %v", in, got, want)
		}
	}
}
