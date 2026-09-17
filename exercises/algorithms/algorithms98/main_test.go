// algorithms98
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two heaps. Implement a structure that accepts numbers one at
// a time and returns the current median.
// Expected asymptotics: O(log n) per addition, O(1) per median query.
package main_test

import (
	"container/heap"
	"testing"
)

type intHeap struct {
	data []int
	less func(a, b int) bool
}

func (h intHeap) Len() int            { return len(h.data) }
func (h intHeap) Less(i, j int) bool  { return h.less(h.data[i], h.data[j]) }
func (h intHeap) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *intHeap) Push(x interface{}) { h.data = append(h.data, x.(int)) }
func (h *intHeap) Pop() interface{} {
	v := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return v
}

type MedianFinder struct {
	low  *intHeap
	high *intHeap
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{
		low:  &intHeap{less: func(a, b int) bool { return a > b }},
		high: &intHeap{less: func(a, b int) bool { return a < b }},
	}
}

func (m *MedianFinder) Add(v int) {
	_ = heap.Init
}

func (m *MedianFinder) Median() float64 {
	return 0
}

func TestMedianFinder(t *testing.T) {
	m := NewMedianFinder()
	m.Add(1)
	if m.Median() != 1 {
		t.Errorf("median after 1 = %v", m.Median())
	}
	m.Add(2)
	if m.Median() != 1.5 {
		t.Errorf("median after 1,2 = %v", m.Median())
	}
	m.Add(3)
	if m.Median() != 2 {
		t.Errorf("median after 1,2,3 = %v", m.Median())
	}
	for _, v := range []int{-10, 100, 50} {
		m.Add(v)
	}
	if m.Median() != 2.5 {
		t.Errorf("median after six values = %v", m.Median())
	}
}
