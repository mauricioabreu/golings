// functions56
// Make the tests pass!

// I AM NOT DONE
//
// minOf must return the smallest argument, or an error when there are none.
// Practices variadic functions and returning (value, error).
package main_test

import (
	"errors"
	"testing"
)

func minOf(nums ...int) (int, error) {
	m := nums[0]
	for _, n := range nums[1:] {
		if n < m {
			m = n
		}
	}
	return m, nil
}

func TestMinOf(t *testing.T) {
	_ = errors.New
	if m, err := minOf(4, -2, 9); err != nil || m != -2 {
		t.Errorf("minOf(4,-2,9) = %d, %v", m, err)
	}
	if _, err := minOf(); err == nil {
		t.Errorf("minOf() should return an error")
	}
}
