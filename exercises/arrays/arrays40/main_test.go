// arrays40
// Make the tests pass!

// I AM NOT DONE
//
// sortScores sorts an array of scores in descending order. Right now the code panics:
// sort.Slice is given an array rather than a slice.
// sort.Slice takes an interface{} and only works on slices.
package main_test

import (
	"sort"
	"testing"
)

func sortScores(a *[5]int) {
	sort.Slice(*a, func(i, j int) bool { return a[i] > a[j] })
}

func TestSortScores(t *testing.T) {
	a := [5]int{3, 9, 1, 7, 5}
	sortScores(&a)
	if a != [5]int{9, 7, 5, 3, 1} {
		t.Errorf("sortScores = %v", a)
	}
}
