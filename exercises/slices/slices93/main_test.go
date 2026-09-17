// slices93
// Make the tests pass!

// I AM NOT DONE
//
// windowMax returns the maximum of every window of size k, by plain scanning.
// Practices the nested slice s[i:i+k] and a search for a maximum.
package main_test

import (
	"reflect"
	"testing"
)

func windowMax(s []int, k int) []int {
	var out []int
	for i := 0; i < len(s)-k; i++ {
		m := 0
		for _, v := range s[i : i+k] {
			if v > m {
				m = v
			}
		}
		out = append(out, m)
	}
	return out
}

func TestWindowMax(t *testing.T) {
	got := windowMax([]int{-1, -3, -2, -5, -4}, 2)
	if !reflect.DeepEqual(got, []int{-1, -2, -2, -4}) {
		t.Errorf("windowMax = %v", got)
	}
}
