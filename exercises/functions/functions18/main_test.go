// functions18
// Make the tests pass!

// I AM NOT DONE
//
// countNodes counts the nodes of a tree with a recursive anonymous function.
// The code does not compile: the function refers to itself before it is declared.
// Practices declaring a function variable before assigning a recursive literal.
package main_test

import "testing"

func countNodes(tree map[string][]string, root string) int {
	count := func(name string) int {
		n := 1
		for _, child := range tree[name] {
			n += count(child)
		}
		return n
	}
	return count(root)
}

func TestCountNodes(t *testing.T) {
	tree := map[string][]string{
		"root": {"a", "b"},
		"a":    {"c"},
		"b":    {"d", "e"},
	}
	if got := countNodes(tree, "root"); got != 6 {
		t.Errorf("countNodes = %d, want 6", got)
	}
	if got := countNodes(tree, "e"); got != 1 {
		t.Errorf("countNodes(leaf) = %d, want 1", got)
	}
}
