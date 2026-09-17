// range38
// Make the tests pass!

// I AM NOT DONE
//
// sumTail adds up every element but the first.
// Practices using the index in the body of a range.
package main_test

import "testing"

func sumTail(s []int) int {
	total := 0
	for i, v := range s {
		if i == len(s)-1 {
			continue
		}
		total += v
	}
	return total
}

func TestSumTail(t *testing.T) {
	if got := sumTail([]int{100, 1, 2, 3}); got != 6 {
		t.Errorf("sumTail = %d, want 6", got)
	}
}
