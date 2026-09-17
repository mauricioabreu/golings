// primitive_types18
// Make the tests pass!

// I AM NOT DONE
//
// addChecked adds two int64 values and reports an overflow.
// Practices detecting overflow from the signs of the operands and the result.
package main_test

import (
	"math"
	"testing"
)

func addChecked(a, b int64) (int64, bool) {
	s := a + b
	if s < a || s < b {
		return s, false
	}
	return s, true
}

func TestAddChecked(t *testing.T) {
	if s, ok := addChecked(2, 3); !ok || s != 5 {
		t.Errorf("addChecked(2, 3) = %d, %v", s, ok)
	}
	if s, ok := addChecked(5, -3); !ok || s != 2 {
		t.Errorf("addChecked(5, -3) = %d, %v", s, ok)
	}
	if _, ok := addChecked(math.MaxInt64, 1); ok {
		t.Errorf("MaxInt64+1 should overflow")
	}
	if _, ok := addChecked(math.MinInt64, -1); ok {
		t.Errorf("MinInt64-1 should overflow")
	}
}
