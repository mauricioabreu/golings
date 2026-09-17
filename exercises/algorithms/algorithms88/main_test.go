// algorithms88
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: linked lists. Reverse a singly linked list: the last node
// must become the head. Return the new head of the list.
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

func reverseList(head *ListNode) *ListNode {
	return head
}

func TestReverseList(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{7}, []int{7}},
		{nil, []int{}},
	}
	for _, c := range cases {
		if got := toSlice(reverseList(build(c.in...))); !reflect.DeepEqual(got, c.want) {
			t.Errorf("reverseList(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
