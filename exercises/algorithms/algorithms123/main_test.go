// algorithms123
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: knapsack dynamic programming. Each item can be taken at most
// once. Return the maximum total value with the given capacity.
// Expected asymptotics: O(n·W) time, O(W) space.
package main_test

import "testing"

func knapsack(weights, values []int, capacity int) int {
	return 0
}

func TestKnapsack(t *testing.T) {
	cases := []struct {
		w, v      []int
		cap, want int
	}{
		{[]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7, 9},
		{[]int{2, 2, 2}, []int{3, 3, 3}, 5, 6},
		{nil, nil, 10, 0},
		{[]int{5}, []int{10}, 4, 0},
		{[]int{1, 1, 1}, []int{5, 4, 3}, 3, 12},
	}
	for _, c := range cases {
		if got := knapsack(c.w, c.v, c.cap); got != c.want {
			t.Errorf("knapsack(%v, %v, %d) = %d, want %d", c.w, c.v, c.cap, got, c.want)
		}
	}
}
