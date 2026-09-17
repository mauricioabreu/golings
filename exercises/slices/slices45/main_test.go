// slices45
// Make the tests pass!

// I AM NOT DONE
//
// groupRuns splits a slice into groups of equal consecutive elements.
// Practices slices of the original slice with computed bounds.
package main_test

import (
	"reflect"
	"testing"
)

func groupRuns(s []int) [][]int {
	var out [][]int
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		out = append(out, s[i:j-1])
		i++
	}
	return out
}

func TestGroupRuns(t *testing.T) {
	got := groupRuns([]int{1, 1, 2, 3, 3, 3})
	if !reflect.DeepEqual(got, [][]int{{1, 1}, {2}, {3, 3, 3}}) {
		t.Errorf("groupRuns = %v", got)
	}
}
