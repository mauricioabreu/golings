// slices34
// Make the tests pass!

// I AM NOT DONE
//
// pairs returns every pair of neighbouring elements.
// Practices the loop bounds when reaching for s[i+1].
package main_test

import (
	"reflect"
	"testing"
)

func pairs(s []int) [][2]int {
	var out [][2]int
	for i := range s {
		out = append(out, [2]int{s[i], s[i]})
	}
	return out
}

func TestPairs(t *testing.T) {
	if got := pairs([]int{1, 2, 3}); !reflect.DeepEqual(got, [][2]int{{1, 2}, {2, 3}}) {
		t.Errorf("pairs = %v", got)
	}
	if got := pairs([]int{1}); len(got) != 0 {
		t.Errorf("pairs([1]) = %v", got)
	}
}
