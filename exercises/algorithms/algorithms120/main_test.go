// algorithms120
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: sorting intervals. Merge overlapping and touching intervals and
// return them in ascending order of start.
// Expected asymptotics: O(n·log n) time, O(n) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func mergeIntervals(intervals [][2]int) [][2]int {
	_ = sort.Ints
	return nil
}

func TestMergeIntervals(t *testing.T) {
	cases := []struct {
		in, want [][2]int
	}{
		{[][2]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][2]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][2]int{{1, 4}, {4, 5}}, [][2]int{{1, 5}}},
		{[][2]int{{5, 6}, {1, 2}}, [][2]int{{1, 2}, {5, 6}}},
		{nil, nil},
		{[][2]int{{1, 10}, {2, 3}, {4, 5}}, [][2]int{{1, 10}}},
	}
	for _, c := range cases {
		if got := mergeIntervals(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("mergeIntervals(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
