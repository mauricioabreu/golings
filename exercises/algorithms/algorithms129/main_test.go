// algorithms129
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: greedy. A cup costs 5; customers pay with bills of 5, 10, or 20.
// Return true if change could be given to everyone (the register starts
// empty).
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func lemonadeChange(bills []int) bool {
	return false
}

func TestLemonadeChange(t *testing.T) {
	cases := []struct {
		bills []int
		want  bool
	}{
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10, 10, 20}, false},
		{[]int{10}, false},
		{nil, true},
		{[]int{5, 5, 5, 5, 5, 5, 20, 20}, true},
		{[]int{5, 5, 5, 5, 20, 20}, false},
	}
	for _, c := range cases {
		if got := lemonadeChange(c.bills); got != c.want {
			t.Errorf("lemonadeChange(%v) = %v, want %v", c.bills, got, c.want)
		}
	}
}
