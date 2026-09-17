// algorithms22
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: binary search on the answer. Return the integer part of the
// square root of a non-negative x without math.Sqrt.
// Expected asymptotics: O(log x) time, O(1) space.
package main_test

import "testing"

func mySqrt(x int) int {
	return 0
}

func TestMySqrt(t *testing.T) {
	cases := map[int]int{0: 0, 1: 1, 4: 2, 8: 2, 15: 3, 16: 4, 2147395599: 46339, 1 << 62: 1 << 31, 9223372036854775807: 3037000499}
	for in, want := range cases {
		if got := mySqrt(in); got != want {
			t.Errorf("mySqrt(%d) = %d, want %d", in, got, want)
		}
	}
}
