// slices40
// Make the tests pass!

// I AM NOT DONE
//
// pop returns the slice without its last element.
// Practices the slice expression s[:high].
package main_test

import (
	"reflect"
	"testing"
)

func pop(s []int) []int {
	return s[:len(s)]
}

func TestPop(t *testing.T) {
	if got := pop([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("pop = %v", got)
	}
}
