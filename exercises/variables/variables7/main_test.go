// variables7
// Make the tests pass!

// I AM NOT DONE
//
// minValue must return the smallest element of a non-empty slice.
// For positive numbers it always returns 0.
// Practices picking the right initial value for a variable.
package main_test

import "testing"

func minValue(nums []int) int {
	var m int
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}

func TestMinValue(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{{[]int{5, 3, 8}, 3}, {[]int{-1, -7}, -7}, {[]int{42}, 42}}
	for _, c := range cases {
		if got := minValue(c.in); got != c.want {
			t.Errorf("minValue(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
