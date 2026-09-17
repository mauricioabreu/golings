// anonymous_functions36
// Make the tests pass!

// I AM NOT DONE
//
// walk traverses a tree; the visitor literal returns false to keep the walk out of
// the children of a node. collect must not descend into branches named "skip".
// Practices a literal steering a traversal.
package main_test

import (
	"reflect"
	"testing"
)

type Node struct {
	Name     string
	Children []*Node
}

func walk(n *Node, visit func(*Node) bool) {
	visit(n)
	for _, c := range n.Children {
		walk(c, visit)
	}
}

func collect(root *Node) []string {
	var names []string
	walk(root, func(n *Node) bool {
		if n.Name == "skip" {
			return false
		}
		names = append(names, n.Name)
		return true
	})
	return names
}

func TestCollect(t *testing.T) {
	root := &Node{"root", []*Node{
		{"a", nil},
		{"skip", []*Node{{"hidden", nil}}},
		{"b", []*Node{{"c", nil}}},
	}}
	if got := collect(root); !reflect.DeepEqual(got, []string{"root", "a", "b", "c"}) {
		t.Errorf("collect = %v", got)
	}
}
