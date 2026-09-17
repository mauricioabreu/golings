// algorithms28
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: min-heap (or binary search with traversal). On a height map
// the water rises over time; at time t you can swim over cells with
// height at most t. Return the minimum time to get from (0,0) to the
// bottom-right corner.
// Expected asymptotics: O(r·c·log(r·c)) time, O(r·c) space.
package main_test

import (
	"container/heap"
	"testing"
)

type wcell struct{ h, r, c int }

type wheap []wcell

func (x wheap) Len() int            { return len(x) }
func (x wheap) Less(i, j int) bool  { return x[i].h < x[j].h }
func (x wheap) Swap(i, j int)       { x[i], x[j] = x[j], x[i] }
func (x *wheap) Push(v interface{}) { *x = append(*x, v.(wcell)) }
func (x *wheap) Pop() interface{} {
	old := *x
	v := old[len(old)-1]
	*x = old[:len(old)-1]
	return v
}

func swimInWater(grid [][]int) int {
	_ = heap.Init
	return 0
}

func TestSwimInWater(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 2}, {1, 3}}, 3},
		{[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, 16},
		{[][]int{{0}}, 0},
		{nil, 0},
	}
	for _, c := range cases {
		if got := swimInWater(c.grid); got != c.want {
			t.Errorf("swimInWater = %d, want %d", got, c.want)
		}
	}
}
