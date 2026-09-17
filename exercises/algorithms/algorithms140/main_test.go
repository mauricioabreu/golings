// algorithms140
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: linked lists with carry. Numbers are stored in lists as digits
// in reverse order. Return the sum in the same format.
// Expected asymptotics: O(max(n, m)) time, O(max(n, m)) space.
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

func addTwoNumbers(a, b *ListNode) *ListNode {
	return nil
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct{ a, b, want []int }{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9}, []int{9, 9}, []int{8, 9, 0, 0, 1}},
		{nil, []int{1}, []int{1}},
	}
	for _, c := range cases {
		if got := toSlice(addTwoNumbers(build(c.a...), build(c.b...))); !reflect.DeepEqual(got, c.want) {
			t.Errorf("add(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
