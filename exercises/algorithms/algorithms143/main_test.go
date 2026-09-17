// algorithms143
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: sorting and a min-heap of end times. Find the minimum number
// of meeting rooms needed to hold all meetings [start, end).
// Expected asymptotics: O(n·log n) time, O(n) space.
package main_test

import (
	"container/heap"
	"sort"
	"testing"
)

type endHeap []int

func (h endHeap) Len() int            { return len(h) }
func (h endHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h endHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *endHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *endHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func minMeetingRooms(intervals [][2]int) int {
	_ = heap.Init
	_ = sort.Ints
	return 0
}

func TestMinMeetingRooms(t *testing.T) {
	cases := []struct {
		in   [][2]int
		want int
	}{
		{[][2]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{[][2]int{{7, 10}, {2, 4}}, 1},
		{nil, 0},
		{[][2]int{{1, 5}, {2, 6}, {3, 7}}, 3},
		{[][2]int{{1, 2}, {2, 3}, {3, 4}}, 1},
	}
	for _, c := range cases {
		if got := minMeetingRooms(c.in); got != c.want {
			t.Errorf("minMeetingRooms(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
