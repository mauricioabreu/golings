// arrays84
// Make the tests pass!

// I AM NOT DONE
//
// sum must add up the elements of an array, and adds up their indexes instead.
// With a single variable a range yields the index, not the value.
package main_test

import "testing"

func sum(a [4]int) int {
	s := 0
	for v := range a {
		s += v
	}
	return s
}

func TestSum(t *testing.T) {
	if got := sum([4]int{10, 20, 30, 40}); got != 100 {
		t.Errorf("sum = %d, want 100", got)
	}
}
