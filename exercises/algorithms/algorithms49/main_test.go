// algorithms49
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: recursion with early exit. A tree is balanced if the depths of
// every node's subtrees differ by no more than 1.
// Expected asymptotics: O(n) time, O(h) space.
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return false
}

func TestIsBalanced(t *testing.T) {
	ok := &TreeNode{3, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}
	if !isBalanced(ok) || !isBalanced(nil) {
		t.Errorf("balanced tree rejected")
	}
	bad := &TreeNode{1,
		&TreeNode{2, &TreeNode{3, &TreeNode{Val: 4}, nil}, nil},
		nil}
	if isBalanced(bad) {
		t.Errorf("unbalanced tree accepted")
	}
}
