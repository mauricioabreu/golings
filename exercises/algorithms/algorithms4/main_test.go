// algorithms4
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: breadth-first traversal. Return the node values level by level:
// the first level is the root, then its children, and so on.
// Expected asymptotics: O(n) time, O(n) space.
package main_test

import (
	"reflect"
	"testing"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	return nil
}

func TestLevelOrder(t *testing.T) {
	tree := &TreeNode{3,
		&TreeNode{Val: 9},
		&TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}
	if got := levelOrder(tree); !reflect.DeepEqual(got, [][]int{{3}, {9, 20}, {15, 7}}) {
		t.Errorf("levelOrder = %v", got)
	}
	if got := levelOrder(nil); len(got) != 0 {
		t.Errorf("levelOrder(nil) = %v", got)
	}
	if got := levelOrder(&TreeNode{Val: 1}); !reflect.DeepEqual(got, [][]int{{1}}) {
		t.Errorf("single node = %v", got)
	}
}
