// algorithms92
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: traversal with markers for empty nodes. Turn a tree into a string
// and back so that the restored tree equals the original.
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func same(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && same(a.Left, b.Left) && same(a.Right, b.Right)
}

func serialize(root *TreeNode) string {
	return ""
}

func deserialize(s string) *TreeNode {
	return nil
}

func TestSerialize(t *testing.T) {
	_, _ = strconv.Itoa, strings.Fields
	trees := []*TreeNode{
		{1, &TreeNode{Val: 2}, &TreeNode{3, &TreeNode{Val: 4}, &TreeNode{Val: 5}}},
		nil,
		{Val: -7},
		{0, &TreeNode{Val: -1, Left: &TreeNode{Val: 2}}, nil},
	}
	for _, tree := range trees {
		if got := deserialize(serialize(tree)); !same(got, tree) {
			t.Errorf("round trip failed for %v", serialize(tree))
		}
	}
}
