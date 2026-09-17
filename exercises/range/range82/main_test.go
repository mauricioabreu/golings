// range82
// Make the tests pass!

// I AM NOT DONE
//
// splitOn splits a slice of numbers into parts on the separator value sep, which is left out of the parts.
// Practices a range building up the current part.
package main_test

import (
	"reflect"
	"testing"
)

func splitOn(s []int, sep int) [][]int {
	var out [][]int
	cur := []int{}
	for _, v := range s {
		if v == sep {
			continue
		}
		cur = append(cur, v)
	}
	return out
}

func TestSplitOn(t *testing.T) {
	got := splitOn([]int{1, 2, 0, 3, 0, 0, 4}, 0)
	want := [][]int{{1, 2}, {3}, {}, {4}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitOn = %v, want %v", got, want)
	}
}
