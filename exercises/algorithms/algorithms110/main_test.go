// algorithms110
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: recursion with a hash table of positions. Given the preorder
// and inorder traversals (values are unique), rebuild the tree.
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func preorderOf(t *TreeNode, out []int) []int {
	if t == nil {
		return out
	}
	out = append(out, t.Val)
	out = preorderOf(t.Left, out)
	return preorderOf(t.Right, out)
}

func inorderOf(t *TreeNode, out []int) []int {
	if t == nil {
		return out
	}
	out = inorderOf(t.Left, out)
	out = append(out, t.Val)
	return inorderOf(t.Right, out)
}

func buildTree(preorder, inorder []int) *TreeNode {
	return nil
}

func TestBuildTree(t *testing.T) {
	trees := []*TreeNode{
		{3, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}},
		{Val: -1},
		nil,
		{1, nil, &TreeNode{2, &TreeNode{Val: 3}, nil}},
	}
	for i, tree := range trees {
		pre := preorderOf(tree, nil)
		in := inorderOf(tree, nil)
		got := buildTree(pre, in)
		if len(preorderOf(got, nil)) != len(pre) {
			t.Fatalf("case %d: sizes differ", i)
		}
		gotPre := preorderOf(got, nil)
		gotIn := inorderOf(got, nil)
		for j := range pre {
			if gotPre[j] != pre[j] || gotIn[j] != in[j] {
				t.Errorf("case %d: traversals differ", i)
				break
			}
		}
	}
}
