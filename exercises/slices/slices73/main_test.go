// slices73
// Make the tests pass!

// I AM NOT DONE
//
// concat joins two slices.
// Practices append(a, b...).
package main_test

import (
	"reflect"
	"testing"
)

func concat(a, b []int) []int {
	return append(a, b)
}

func TestConcat(t *testing.T) {
	if got := concat([]int{1, 2}, []int{3, 4}); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("concat = %v", got)
	}
}
