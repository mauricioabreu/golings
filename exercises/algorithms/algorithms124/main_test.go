// algorithms124
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: sorting intervals. Determine whether all meetings [start, end)
// can be attended, i.e. whether there are no overlaps.
// Expected asymptotics: O(n·log n) time, O(n) space.
package main_test

import (
	"sort"
	"testing"
)

func canAttendMeetings(intervals [][2]int) bool {
	_ = sort.Ints
	return false
}

func TestCanAttendMeetings(t *testing.T) {
	cases := []struct {
		in   [][2]int
		want bool
	}{
		{[][2]int{{0, 30}, {5, 10}, {15, 20}}, false},
		{[][2]int{{7, 10}, {2, 4}}, true},
		{nil, true},
		{[][2]int{{1, 2}, {2, 3}}, true},
		{[][2]int{{1, 5}, {4, 6}}, false},
	}
	for _, c := range cases {
		if got := canAttendMeetings(c.in); got != c.want {
			t.Errorf("canAttendMeetings(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
