// algorithms112
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: reversing sublists. Reverse the nodes of a list in groups of k;
// an incomplete final group stays as is.
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	return head
}

func TestReverseKGroup(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3}, 1, []int{1, 2, 3}},
		{[]int{1, 2}, 3, []int{1, 2}},
		{nil, 2, []int{}},
	}
	for _, c := range cases {
		if got := toSlice(reverseKGroup(build(c.in...), c.k)); !reflect.DeepEqual(got, c.want) {
			t.Errorf("reverseKGroup(%v, %d) = %v, want %v", c.in, c.k, got, c.want)
		}
	}
}
