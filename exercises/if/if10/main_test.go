// if10
// Make the tests pass!

// I AM NOT DONE
//
// sign must return -1, 0 or 1 depending on the sign of the number.
// Practices an if / else if / else chain.
package main_test

import "testing"

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x == 0 {
		return 0
	} else {
		return 1
	}
}

func TestSign(t *testing.T) {
	cases := map[int]int{5: 1, 0: 0, -3: -1}
	for in, want := range cases {
		if got := sign(in); got != want {
			t.Errorf("sign(%d) = %d, want %d", in, got, want)
		}
	}
}
