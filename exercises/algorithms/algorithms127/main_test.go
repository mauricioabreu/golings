// algorithms127
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a heap. Merge k sorted lists into a single sorted list and
// return its head. There may be no lists at all.
// Expected asymptotics: O(N·log k) time, O(k) space.
package main_test

import (
	"container/heap"
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func build(vals ...int) *ListNode {
	var head *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		head = &ListNode{vals[i], head}
	}
	return head
}

func toSlice(l *ListNode) []int {
	out := []int{}
	for ; l != nil; l = l.Next {
		out = append(out, l.Val)
	}
	return out
}

type nodeHeap []*ListNode

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := old[len(old)-1]
	*h = old[:len(old)-1]
	return n
}

func mergeKLists(lists []*ListNode) *ListNode {
	_ = heap.Init
	return nil
}

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{build(1, 4, 5), build(1, 3, 4), build(2, 6), nil}
	if got := toSlice(mergeKLists(lists)); !reflect.DeepEqual(got, []int{1, 1, 2, 3, 4, 4, 5, 6}) {
		t.Errorf("mergeKLists = %v", got)
	}
	if got := toSlice(mergeKLists(nil)); len(got) != 0 {
		t.Errorf("mergeKLists(nil) = %v", got)
	}
	if got := toSlice(mergeKLists([]*ListNode{build(3)})); !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("single list = %v", got)
	}
}
