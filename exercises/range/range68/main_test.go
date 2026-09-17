// range68
// Make the tests pass!

// I AM NOT DONE
//
// runningMax returns, for every position, the maximum of the elements up to and including it.
// Practices an accumulator in a range plus writing by index.
package main_test

import (
	"reflect"
	"testing"
)

func runningMax(s []int) []int {
	out := make([]int, len(s))
	m := 0
	for i, v := range s {
		if v > m {
			out[i] = v
		}
	}
	return out
}

func TestRunningMax(t *testing.T) {
	if got := runningMax([]int{-3, -5, 2, 1, 4}); !reflect.DeepEqual(got, []int{-3, -3, 2, 2, 4}) {
		t.Errorf("runningMax = %v", got)
	}
}
