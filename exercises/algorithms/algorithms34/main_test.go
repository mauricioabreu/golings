// algorithms34
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two-pass greedy. Each child needs at least one candy, and a
// child with a higher rating gets more candies than neighbors with a
// lower rating. Return the minimum number of candies.
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import "testing"

func candy(ratings []int) int {
	return 0
}

func TestCandy(t *testing.T) {
	cases := []struct {
		ratings []int
		want    int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{nil, 0},
		{[]int{7}, 1},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{5, 4, 3, 2, 1}, 15},
	}
	for _, c := range cases {
		if got := candy(c.ratings); got != c.want {
			t.Errorf("candy(%v) = %d, want %d", c.ratings, got, c.want)
		}
	}
}
