// functions82
// Make the tests pass!

// I AM NOT DONE
//
// sum takes a variable number of arguments.
// A whole slice has to be handed to it. The code does not compile.
// Practices passing a slice to a variadic function.
package main_test

import "testing"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func sumSlice(values []int) int {
	return sum(values)
}

func TestSumSlice(t *testing.T) {
	if got := sumSlice([]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("sumSlice(1,2,3,4) = %d, want 10", got)
	}
}
