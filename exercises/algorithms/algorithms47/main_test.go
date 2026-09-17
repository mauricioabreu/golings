// algorithms47
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: cycle detection (Floyd). A number is happy if repeatedly
// replacing it with the sum of the squares of its digits leads to 1.
// Expected asymptotics: O(log n) time, O(1) space.
package main_test

import "testing"

func isHappy(n int) bool {
	return false
}

func TestIsHappy(t *testing.T) {
	cases := map[int]bool{19: true, 2: false, 1: true, 7: true, 116: false}
	for in, want := range cases {
		if got := isHappy(in); got != want {
			t.Errorf("isHappy(%d) = %v, want %v", in, got, want)
		}
	}
}
