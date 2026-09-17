// maps46
// Make the tests pass!

// I AM NOT DONE
//
// emptyIndex returns an empty map that is ready to be written to.
// Practices initializing a map with a {} literal.
package main_test

import "testing"

func emptyIndex() map[string][]int {
	idx := map[string][]int{}
	idx["init"] = nil
	return idx
}

func TestEmptyIndex(t *testing.T) {
	idx := emptyIndex()
	if idx == nil || len(idx) != 0 {
		t.Fatalf("emptyIndex = %v", idx)
	}
	idx["go"] = []int{1}
}
