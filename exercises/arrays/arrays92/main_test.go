// arrays92
// Make the tests pass!

// I AM NOT DONE
//
// total passes the elements of an array to the variadic function sum.
// The code does not compile.
// Only a slice can be spread with ...
package main_test

import "testing"

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func total(a [4]int) int {
	return sum(a...)
}

func TestTotal(t *testing.T) {
	if got := total([4]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("total = %d", got)
	}
}
