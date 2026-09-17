// algorithms54
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: max-heap. While more than one stone remains, take the two
// heaviest: equal ones both disappear, unequal ones yield a stone with the
// difference of their weights. Return the weight of the last stone or 0.
// Expected asymptotics: O(n·log n) time, O(n) space.
package main_test

import (
	"container/heap"
	"testing"
)

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func lastStoneWeight(stones []int) int {
	_ = heap.Init
	return 0
}

func TestLastStoneWeight(t *testing.T) {
	cases := []struct {
		stones []int
		want   int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1}, 1},
		{[]int{2, 2}, 0},
		{nil, 0},
		{[]int{10, 4, 2, 10}, 2},
	}
	for _, c := range cases {
		if got := lastStoneWeight(c.stones); got != c.want {
			t.Errorf("lastStoneWeight(%v) = %d, want %d", c.stones, got, c.want)
		}
	}
}
