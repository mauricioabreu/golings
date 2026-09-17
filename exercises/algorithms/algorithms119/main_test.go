// algorithms119
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: linked lists and a dummy head. Merge two sorted lists into one
// sorted list, reusing the nodes.
// Expected asymptotics: O(n + m) time, O(1) space.
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

func mergeTwoLists(a, b *ListNode) *ListNode {
	return a
}

func TestMergeTwoLists(t *testing.T) {
	cases := []struct{ a, b, want []int }{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{nil, nil, []int{}},
		{nil, []int{0}, []int{0}},
		{[]int{5}, []int{1, 2, 3}, []int{1, 2, 3, 5}},
	}
	for _, c := range cases {
		if got := toSlice(mergeTwoLists(build(c.a...), build(c.b...))); !reflect.DeepEqual(got, c.want) {
			t.Errorf("merge(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
