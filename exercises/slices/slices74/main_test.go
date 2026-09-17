// slices74
// Make the tests pass!

// I AM NOT DONE
//
// runningSum returns the running sums without changing the input.
// Practices a new slice of the same length plus an accumulator.
package main_test

import (
	"reflect"
	"testing"
)

func runningSum(s []int) []int {
	for i := 1; i < len(s); i++ {
		s[i] += s[i-1]
	}
	return s
}

func TestRunningSum(t *testing.T) {
	in := []int{1, 2, 3, 4}
	if got := runningSum(in); !reflect.DeepEqual(got, []int{1, 3, 6, 10}) {
		t.Errorf("runningSum = %v", got)
	}
	if !reflect.DeepEqual(in, []int{1, 2, 3, 4}) {
		t.Errorf("input modified: %v", in)
	}
}
