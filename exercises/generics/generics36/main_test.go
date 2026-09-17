// generics36
// Make the tests pass!

// I AM NOT DONE
//
// Set[[]int] does not compile: slices do not satisfy comparable.
// Distinct sets of numbers have to be stored.
// The comparable constraint rules out slices, maps and functions.
package main_test

import (
	"fmt"
	"testing"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) { s[v] = struct{}{} }

func uniqueCombos(combos [][]int) int {
	s := Set[[]int]{}
	for _, c := range combos {
		s.Add(c)
	}
	_ = fmt.Sprint
	return len(s)
}

func TestUniqueCombos(t *testing.T) {
	if got := uniqueCombos([][]int{{1, 2}, {2, 1}, {1, 2}}); got != 2 {
		t.Errorf("uniqueCombos = %d, want 2", got)
	}
}
