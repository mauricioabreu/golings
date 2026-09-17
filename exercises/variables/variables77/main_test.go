// variables77
// Make the tests pass!

// I AM NOT DONE
//
// reverseDigits must write the digits of a non-negative number in reverse order.
// Practices integer division and remainders while updating several variables.
package main_test

import "testing"

func reverseDigits(n int) int {
	result := 0
	for n > 0 {
		result = result + n%10
		n = n % 10
	}
	return result
}

func TestReverseDigits(t *testing.T) {
	cases := map[int]int{123: 321, 1200: 21, 7: 7, 0: 0}
	for in, want := range cases {
		if got := reverseDigits(in); got != want {
			t.Errorf("reverseDigits(%d) = %d, want %d", in, got, want)
		}
	}
}
