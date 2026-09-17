// algorithms68
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: working with a number as a digit array. The digits are stored in
// a slice from most significant to least significant. Add 1 and return the
// result.
// Expected asymptotics: O(n) time, O(1) extra space.
package main_test

import (
	"reflect"
	"testing"
)

func plusOne(digits []int) []int {
	return nil
}

func TestPlusOne(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 9}, []int{4, 4, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{0}, []int{1}},
		{nil, []int{1}},
	}
	for _, c := range cases {
		if got := plusOne(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("plusOne(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
