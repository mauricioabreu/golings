// maps101
// Make the tests pass!

// I AM NOT DONE
//
// clearAll removes every element while keeping the map itself, as others refer to it.
// A delete during a range is safe.
package main_test

import "testing"

func clearAll(m map[string]int) {
	for k := range m {
		m[k] = 0
	}
}

func TestClearAll(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	alias := m
	clearAll(m)
	if len(alias) != 0 {
		t.Errorf("alias = %v, want empty", alias)
	}
}
