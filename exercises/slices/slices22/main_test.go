// slices22
// Make the tests pass!

// I AM NOT DONE
//
// evens returns the even numbers only.
// Practices an append inside a condition.
package main_test

import (
	"reflect"
	"testing"
)

func evens(s []int) []int {
	var out []int
	for _, v := range s {
		if v%2 == 0 {
			out = append(out, v%2)
		}
	}
	return out
}

func TestEvens(t *testing.T) {
	if got := evens([]int{1, 2, 3, 4, 6}); !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("evens = %v", got)
	}
}
