// algorithms62
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: tree recursion. Swap the left and right subtrees at every level
// and return the root.
// Expected asymptotics: O(n) time, O(h) space.
package main_test

import (
	"reflect"
	"testing"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorder(t *TreeNode, out []int) []int {
	if t == nil {
		return out
	}
	out = inorder(t.Left, out)
	out = append(out, t.Val)
	return inorder(t.Right, out)
}

func invertTree(root *TreeNode) *TreeNode {
	return root
}

func TestInvertTree(t *testing.T) {
	tree := &TreeNode{4,
		&TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
		&TreeNode{7, &TreeNode{Val: 6}, &TreeNode{Val: 9}}}
	got := inorder(invertTree(tree), nil)
	if !reflect.DeepEqual(got, []int{9, 7, 6, 4, 3, 2, 1}) {
		t.Errorf("inverted inorder = %v", got)
	}
	if invertTree(nil) != nil {
		t.Errorf("invertTree(nil) should be nil")
	}
}
