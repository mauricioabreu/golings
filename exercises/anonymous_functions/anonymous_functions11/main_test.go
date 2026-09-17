// anonymous_functions11
// Make the tests pass!

// I AM NOT DONE
//
// The literal has to build up a running total, and the total stays at zero.
// := in the body of a literal declares a new local variable.
package main_test

import "testing"

func sumAll(groups [][]int) int {
	total := 0
	add := func(vs []int) {
		for _, v := range vs {
			total := total + v
			_ = total
		}
	}
	for _, g := range groups {
		add(g)
	}
	return total
}

func TestSumAll(t *testing.T) {
	if got := sumAll([][]int{{1, 2}, {3}}); got != 6 {
		t.Errorf("sumAll = %d, want 6", got)
	}
}
