// functions43
// Make the tests pass!

// I AM NOT DONE
//
// digitSum must compute the digit sum of a non-negative number recursively.
// Practices recursion and integer arithmetic.
package main_test

import "testing"

func digitSum(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + digitSum(n%10)
}

func TestDigitSum(t *testing.T) {
	cases := map[int]int{0: 0, 7: 7, 123: 6, 9999: 36}
	for in, want := range cases {
		if got := digitSum(in); got != want {
			t.Errorf("digitSum(%d) = %d, want %d", in, got, want)
		}
	}
}
