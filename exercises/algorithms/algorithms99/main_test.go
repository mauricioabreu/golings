// algorithms99
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: recursion with a side result. Return the length of the longest
// path between any two nodes (in edges).
// Expected asymptotics: O(n) time, O(h) space.
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func diameter(root *TreeNode) int {
	return 0
}

func TestDiameter(t *testing.T) {
	tree := &TreeNode{1,
		&TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 5}},
		&TreeNode{Val: 3}}
	if got := diameter(tree); got != 3 {
		t.Errorf("diameter = %d, want 3", got)
	}
	if diameter(nil) != 0 || diameter(&TreeNode{Val: 1}) != 0 {
		t.Errorf("edge cases failed")
	}
	chain := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	if got := diameter(chain); got != 2 {
		t.Errorf("chain diameter = %d, want 2", got)
	}
}
