// algorithms107
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: two pointers with a lag. Remove the n-th node from the end in
// one pass (1 <= n <= length of the list) and return the head.
// Expected asymptotics: O(n) time, O(1) space.
package main_test

import (
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		in   []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
	}
	for _, c := range cases {
		if got := toSlice(removeNthFromEnd(build(c.in...), c.n)); !reflect.DeepEqual(got, c.want) {
			t.Errorf("removeNthFromEnd(%v, %d) = %v, want %v", c.in, c.n, got, c.want)
		}
	}
}
