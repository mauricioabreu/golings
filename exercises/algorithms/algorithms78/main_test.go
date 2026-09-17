// algorithms78
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: recursion with a side result. A path is a sequence of nodes
// connected by edges, going in any direction but without repeats. Find the
// maximum sum of values along such a path (the tree is non-empty, the
// values are arbitrary).
// Expected asymptotics: O(n) time, O(h) space.
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	return 0
}

func TestMaxPathSum(t *testing.T) {
	cases := []struct {
		tree *TreeNode
		want int
	}{
		{&TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}, 6},
		{&TreeNode{-10, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}, 42},
		{&TreeNode{Val: -3}, -3},
		{&TreeNode{-2, &TreeNode{Val: -1}, nil}, -1},
	}
	for i, c := range cases {
		if got := maxPathSum(c.tree); got != c.want {
			t.Errorf("case %d: maxPathSum = %d, want %d", i, got, c.want)
		}
	}
}
