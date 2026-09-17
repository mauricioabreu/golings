// algorithms39
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: parallel traversal of two subtrees. Check that the tree is
// symmetric about its center.
// Expected asymptotics: O(n) time, O(h) space.
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return false
}

func TestIsSymmetric(t *testing.T) {
	sym := &TreeNode{1,
		&TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 4}},
		&TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 3}}}
	if !isSymmetric(sym) || !isSymmetric(nil) || !isSymmetric(&TreeNode{Val: 1}) {
		t.Errorf("symmetric trees rejected")
	}
	asym := &TreeNode{1,
		&TreeNode{2, nil, &TreeNode{Val: 3}},
		&TreeNode{2, nil, &TreeNode{Val: 3}}}
	if isSymmetric(asym) {
		t.Errorf("asymmetric tree accepted")
	}
}
