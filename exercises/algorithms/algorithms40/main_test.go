// algorithms40
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: recursion with bounds. Check that the tree is a valid binary
// search tree: strictly smaller on the left, strictly greater on the right.
// Expected asymptotics: O(n) time, O(h) space.
package main_test

import (
	"math"
	"testing"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return false
}

func TestIsValidBST(t *testing.T) {
	_ = math.MaxInt64
	valid := &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
	if !isValidBST(valid) || !isValidBST(nil) {
		t.Errorf("valid BST rejected")
	}
	invalid := &TreeNode{5, &TreeNode{Val: 1}, &TreeNode{4, &TreeNode{Val: 3}, &TreeNode{Val: 6}}}
	if isValidBST(invalid) {
		t.Errorf("invalid BST accepted")
	}
	dup := &TreeNode{2, &TreeNode{Val: 2}, nil}
	if isValidBST(dup) {
		t.Errorf("duplicate values must be rejected")
	}
}
