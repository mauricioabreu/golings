// algorithms148
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: in-order traversal of a BST. Return the k-th smallest value
// (k is 1-based). If there are fewer than k nodes — return -1.
// Expected asymptotics: O(h + k) time, O(h) space.
package main_test

import "testing"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	return 0
}

func TestKthSmallest(t *testing.T) {
	tree := &TreeNode{5,
		&TreeNode{3, &TreeNode{Val: 2}, &TreeNode{Val: 4}},
		&TreeNode{Val: 6}}
	cases := map[int]int{1: 2, 3: 4, 5: 6, 6: -1}
	for k, want := range cases {
		if got := kthSmallest(tree, k); got != want {
			t.Errorf("kthSmallest(%d) = %d, want %d", k, got, want)
		}
	}
	if kthSmallest(nil, 1) != -1 {
		t.Errorf("empty tree should give -1")
	}
}
