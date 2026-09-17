// algorithms137
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: Boyer–Moore voting. It is guaranteed that a non-empty slice
// contains an element occurring more than n/2 times. Find it without extra
// memory.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import "testing"

func majorityElement(nums []int) int {
	return 0
}

func TestMajorityElement(t *testing.T) {
	big := make([]int, 100001)
	for i := range big {
		if i%2 == 0 {
			big[i] = -7
		} else {
			big[i] = i
		}
	}
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{9}, 9},
		{big, -7},
	}
	for _, c := range cases {
		if got := majorityElement(c.nums); got != c.want {
			t.Errorf("majorityElement(len=%d) = %d, want %d", len(c.nums), got, c.want)
		}
	}
}
