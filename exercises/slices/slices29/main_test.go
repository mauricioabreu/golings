// slices29
// Make the tests pass!

// I AM NOT DONE
//
// padded copies the data into a slice of length n, leaving the rest at zero.
// Practices copy into a slice built in advance.
package main_test

import (
	"reflect"
	"testing"
)

func padded(data []int, n int) []int {
	out := make([]int, n)
	copy(data, out)
	return out
}

func TestPadded(t *testing.T) {
	if got := padded([]int{1, 2}, 4); !reflect.DeepEqual(got, []int{1, 2, 0, 0}) {
		t.Errorf("padded = %v", got)
	}
}
