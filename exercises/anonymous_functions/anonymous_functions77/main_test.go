// anonymous_functions77
// Make the tests pass!

// I AM NOT DONE
//
// apply applies a literal to every element of a matrix, passing the coordinates in,
// and returns a new matrix.
// Practices a literal with several parameters plus new slices.
package main_test

import (
	"reflect"
	"testing"
)

func apply(m [][]int, f func(v, r, c int) int) [][]int {
	for r := range m {
		for c := range m[r] {
			m[r][c] = f(m[r][c], c, r)
		}
	}
	return m
}

func TestApply(t *testing.T) {
	m := [][]int{{1, 2}, {3, 4}, {5, 6}}
	got := apply(m, func(v, r, c int) int { return v * (r + 1) })
	if !reflect.DeepEqual(got, [][]int{{1, 2}, {6, 8}, {15, 18}}) {
		t.Errorf("apply = %v", got)
	}
	if m[1][0] != 3 {
		t.Errorf("input modified: %v", m)
	}
}
