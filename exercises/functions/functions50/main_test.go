// functions50
// Make the tests pass!

// I AM NOT DONE
//
// topScores must return a copy sorted in descending order without changing the original slice.
// The result comes out shuffled.
// The comparison function has to look at the very slice that is being sorted.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func topScores(scores []int) []int {
	sorted := make([]int, len(scores))
	copy(sorted, scores)
	sort.Slice(sorted, func(i, j int) bool { return scores[i] > scores[j] })
	return sorted
}

func TestTopScores(t *testing.T) {
	in := []int{3, 9, 1, 7, 5}
	if got := topScores(in); !reflect.DeepEqual(got, []int{9, 7, 5, 3, 1}) {
		t.Errorf("topScores = %v, want [9 7 5 3 1]", got)
	}
	if !reflect.DeepEqual(in, []int{3, 9, 1, 7, 5}) {
		t.Errorf("input modified: %v", in)
	}
}
