// functions7
// Make the tests pass!

// I AM NOT DONE
//
// product multiplies all of its arguments; with no arguments the result is 1.
// Right now the product is always 0.
// Practices variadic functions and the initial value of an accumulator.
package main_test

import "testing"

func product(nums ...int) int {
	result := 0
	for _, n := range nums {
		result *= n
	}
	return result
}

func TestProduct(t *testing.T) {
	if got := product(2, 3, 4); got != 24 {
		t.Errorf("product(2,3,4) = %d, want 24", got)
	}
	if got := product(); got != 1 {
		t.Errorf("product() = %d, want 1", got)
	}
}
