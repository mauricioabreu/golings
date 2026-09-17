// algorithms93
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: parallel traversal. Check that two trees are identical
// in structure and values.
// Expected asymptotics: O(n) time, O(h) space.
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSameTree(a, b *TreeNode) bool {
	return false
}

func TestIsSameTree(t *testing.T) {
	mk := func() *TreeNode {
		return &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	}
	if !isSameTree(mk(), mk()) || !isSameTree(nil, nil) {
		t.Errorf("equal trees reported as different")
	}
	if isSameTree(mk(), &TreeNode{1, &TreeNode{Val: 2}, nil}) {
		t.Errorf("different structure reported as equal")
	}
	if isSameTree(mk(), &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 4}}) {
		t.Errorf("different values reported as equal")
	}
	if isSameTree(nil, &TreeNode{}) {
		t.Errorf("nil and non-nil are not equal")
	}
}
