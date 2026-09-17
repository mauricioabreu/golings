// range26
// Make the tests pass!

// I AM NOT DONE
//
// countLeaves counts every string value in a nested map[string]interface{}
// structure, at any depth.
// Practices a range with recursion and a type switch.
package main_test

import "testing"

func countLeaves(m map[string]interface{}) int {
	n := 0
	for _, v := range m {
		switch x := v.(type) {
		case string:
			n++
		case map[string]interface{}:
			n++
			_ = x
		}
	}
	return n
}

func TestCountLeaves(t *testing.T) {
	doc := map[string]interface{}{
		"a": "1",
		"b": map[string]interface{}{"c": "2", "d": map[string]interface{}{"e": "3", "f": "4"}},
		"g": 5,
	}
	if got := countLeaves(doc); got != 4 {
		t.Errorf("countLeaves = %d, want 4", got)
	}
}
