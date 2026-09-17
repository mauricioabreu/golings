// algorithms11
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: sorting by capital and a max-heap of profits. Starting with
// capital w, at most k projects may be completed; each requires
// capital[i] and yields profits[i]. Return the maximum final capital.
// Expected asymptotics: O(n·log n) time, O(n) space.
package main_test

import (
	"container/heap"
	"sort"
	"testing"
)

type profitHeap []int

func (h profitHeap) Len() int            { return len(h) }
func (h profitHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h profitHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *profitHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *profitHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func findMaximizedCapital(k, w int, profits, capital []int) int {
	_ = heap.Init
	_ = sort.Ints
	return 0
}

func TestFindMaximizedCapital(t *testing.T) {
	cases := []struct {
		k, w             int
		profits, capital []int
		want             int
	}{
		{2, 0, []int{1, 2, 3}, []int{0, 1, 1}, 4},
		{3, 0, []int{1, 2, 3}, []int{0, 1, 2}, 6},
		{1, 0, []int{5}, []int{5}, 0},
		{0, 7, []int{1}, []int{0}, 7},
		{5, 1, nil, nil, 1},
	}
	for _, c := range cases {
		if got := findMaximizedCapital(c.k, c.w, c.profits, c.capital); got != c.want {
			t.Errorf("findMaximizedCapital(%d, %d, %v, %v) = %d, want %d", c.k, c.w, c.profits, c.capital, got, c.want)
		}
	}
}
