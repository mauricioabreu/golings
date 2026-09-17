// slices70
// Make the tests pass!

// I AM NOT DONE
//
// squares must return the squares of the numbers; the result starts with stray zeros.
// make([]T, n) already holds n elements, and append adds after them.
package main_test

import (
	"reflect"
	"testing"
)

func squares(nums []int) []int {
	out := make([]int, len(nums))
	for _, n := range nums {
		out = append(out, n*n)
	}
	return out
}

func TestSquares(t *testing.T) {
	if got := squares([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{1, 4, 9}) {
		t.Errorf("squares = %v", got)
	}
}
