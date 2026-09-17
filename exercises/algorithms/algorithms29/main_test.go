// algorithms29
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: tree recursion. Return the number of nodes on the longest path
// from the root to a leaf. An empty tree has depth 0.
// Expected asymptotics: O(n) time, O(h) space.
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return 0
}

func TestMaxDepth(t *testing.T) {
	tree := &TreeNode{3,
		&TreeNode{Val: 9},
		&TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}
	if got := maxDepth(tree); got != 3 {
		t.Errorf("maxDepth = %d, want 3", got)
	}
	if maxDepth(nil) != 0 || maxDepth(&TreeNode{Val: 1}) != 1 {
		t.Errorf("edge cases failed")
	}
	deep := &TreeNode{Val: 0}
	cur := deep
	for i := 0; i < 999; i++ {
		cur.Left = &TreeNode{Val: i}
		cur = cur.Left
	}
	if got := maxDepth(deep); got != 1000 {
		t.Errorf("deep tree depth = %d, want 1000", got)
	}
}
