// algorithms91
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a min-heap of size k. Return the k-th largest number of the slice
// (counting duplicates) without sorting the entire slice.
// Expected asymptotics: O(n·log k) time, O(k) space.
package main_test

import (
	"container/heap"
	"testing"
)

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func findKthLargest(nums []int, k int) int {
	_ = heap.Init
	return 0
}

func TestFindKthLargest(t *testing.T) {
	cases := []struct {
		nums    []int
		k, want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{1}, 1, 1},
		{[]int{-1, -2, -3}, 3, -3},
	}
	for _, c := range cases {
		if got := findKthLargest(c.nums, c.k); got != c.want {
			t.Errorf("findKthLargest(%v, %d) = %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
