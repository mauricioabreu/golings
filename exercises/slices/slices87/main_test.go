// slices87
// Make the tests pass!

// I AM NOT DONE
//
// take returns the first n elements, n being no larger than the length.
// Practices s[:n].
package main_test

import (
	"reflect"
	"testing"
)

func take(s []int, n int) []int {
	return s[n:]
}

func TestTake(t *testing.T) {
	if got := take([]int{1, 2, 3, 4}, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("take = %v", got)
	}
}
