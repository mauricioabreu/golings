// anonymous_functions12
// Make the tests pass!

// I AM NOT DONE
//
// each calls a callback for every value, and sum adds them up with a literal.
// Practices a literal changing a captured variable.
package main_test

import "testing"

func each(s []int, f func(int)) {
	for _, v := range s {
		f(v)
	}
}

func sum(s []int) int {
	total := 0
	each(s, func(v int) { total = v })
	return total
}

func TestSum(t *testing.T) {
	if sum([]int{1, 2, 3}) != 6 {
		t.Errorf("sum = %d", sum([]int{1, 2, 3}))
	}
}
