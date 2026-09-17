// functions58
// Make the tests pass!

// I AM NOT DONE
//
// walk traverses the tree in preorder and calls visit for every node.
// Right now it skips the right subtrees and visits a node after its children.
// Practices callback functions and recursion over a tree.
package main_test

import (
	"reflect"
	"testing"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func walk(n *Node, visit func(int)) {
	if n == nil {
		return
	}
	walk(n.Left, visit)
	visit(n.Val)
}

func TestWalk(t *testing.T) {
	root := &Node{1, &Node{2, &Node{4, nil, nil}, nil}, &Node{3, nil, &Node{5, nil, nil}}}
	var got []int
	walk(root, func(v int) { got = append(got, v) })
	if want := []int{1, 2, 4, 3, 5}; !reflect.DeepEqual(got, want) {
		t.Errorf("walk = %v, want %v", got, want)
	}
}
