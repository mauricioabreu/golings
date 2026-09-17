// slices24
// Make the tests pass!

// I AM NOT DONE
//
// chunk splits a slice into parts of size size, the last one possibly shorter.
// Practices slice expressions with computed bounds.
package main_test

import (
	"reflect"
	"testing"
)

func chunk(s []int, size int) [][]int {
	var out [][]int
	for i := 0; i < len(s); i += size {
		out = append(out, s[i:i+size])
	}
	return out
}

func TestChunk(t *testing.T) {
	got := chunk([]int{1, 2, 3, 4, 5}, 2)
	if !reflect.DeepEqual(got, [][]int{{1, 2}, {3, 4}, {5}}) {
		t.Errorf("chunk = %v", got)
	}
	if got := chunk(nil, 3); len(got) != 0 {
		t.Errorf("chunk(nil) = %v", got)
	}
}
