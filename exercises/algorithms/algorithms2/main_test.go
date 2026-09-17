// algorithms2
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a min-heap on the boundary. Compute the volume of water
// trapped on a rectangular height map after rain.
// Expected asymptotics: O(r·c·log(r·c)) time, O(r·c) space.
package main_test

import (
	"container/heap"
	"testing"
)

type cell struct{ h, r, c int }

type cellHeap []cell

func (x cellHeap) Len() int            { return len(x) }
func (x cellHeap) Less(i, j int) bool  { return x[i].h < x[j].h }
func (x cellHeap) Swap(i, j int)       { x[i], x[j] = x[j], x[i] }
func (x *cellHeap) Push(v interface{}) { *x = append(*x, v.(cell)) }
func (x *cellHeap) Pop() interface{} {
	old := *x
	v := old[len(old)-1]
	*x = old[:len(old)-1]
	return v
}

func trapRainWater(m [][]int) int {
	_ = heap.Init
	return 0
}

func TestTrapRainWater(t *testing.T) {
	m := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	if got := trapRainWater(m); got != 4 {
		t.Errorf("trapRainWater = %d, want 4", got)
	}
	m2 := [][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 1, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3},
	}
	if got := trapRainWater(m2); got != 10 {
		t.Errorf("bowl = %d, want 10", got)
	}
	if trapRainWater(nil) != 0 || trapRainWater([][]int{{1, 2}, {3, 4}}) != 0 {
		t.Errorf("small maps should hold no water")
	}
}
