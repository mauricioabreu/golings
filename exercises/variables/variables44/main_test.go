// variables44
// Make the tests pass!

// I AM NOT DONE
//
// This function must return 2 to the power of n by doubling the result in a loop.
// Practices the compound assignment operators (+=, *= and friends).
package main_test

import "testing"

func powerOfTwo(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result += 2
	}
	return result
}

func TestPowerOfTwo(t *testing.T) {
	cases := map[int]int{0: 1, 1: 2, 5: 32, 10: 1024}
	for n, want := range cases {
		if got := powerOfTwo(n); got != want {
			t.Errorf("powerOfTwo(%d) = %d, want %d", n, got, want)
		}
	}
}
