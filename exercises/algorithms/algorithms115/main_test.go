// algorithms115
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: bit arithmetic. Add two integers without using the + and -
// operators. The numbers may be negative.
// Expected asymptotics: O(1) time, O(1) space.
package main_test

import "testing"

func getSum(a, b int) int {
	return 0
}

func TestGetSum(t *testing.T) {
	cases := [][3]int{{1, 2, 3}, {2, 3, 5}, {-1, 1, 0}, {-5, -7, -12}, {0, 0, 0}, {123456789, 987654321, 1111111110}, {-100, 50, -50}}
	for _, c := range cases {
		if got := getSum(c[0], c[1]); got != c[2] {
			t.Errorf("getSum(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
