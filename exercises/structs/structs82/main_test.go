// structs82
// Make the tests pass!

// I AM NOT DONE
//
// Tree.Insert adds a value to a binary search tree, and InOrder returns
// the values in ascending order.
// Practices recursive structs with pointers.
package main_test

import (
	"reflect"
	"testing"
)

type Tree struct {
	Val         int
	Left, Right *Tree
}

func (t *Tree) Insert(v int) *Tree {
	if t == nil {
		return &Tree{Val: v}
	}
	if v < t.Val {
		t.Right = t.Right.Insert(v)
	} else {
		t.Left = t.Left.Insert(v)
	}
	return t
}

func (t *Tree) InOrder(out []int) []int {
	if t == nil {
		return out
	}
	out = t.Left.InOrder(out)
	out = append(out, t.Val)
	return t.Right.InOrder(out)
}

func TestTree(t *testing.T) {
	var root *Tree
	for _, v := range []int{5, 2, 8, 1, 9, 3} {
		root = root.Insert(v)
	}
	if got := root.InOrder(nil); !reflect.DeepEqual(got, []int{1, 2, 3, 5, 8, 9}) {
		t.Errorf("InOrder = %v", got)
	}
}
