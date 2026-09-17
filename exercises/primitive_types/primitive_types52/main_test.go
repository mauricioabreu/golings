// primitive_types52
// Make the tests pass!

// I AM NOT DONE
//
// maxByte must return the largest value of the uint8 type.
// Practices the range of the unsigned types and the math constants.
package main_test

import (
	"math"
	"testing"
)

func maxByte() uint8 {
	return math.MaxInt8
}

func TestMaxByte(t *testing.T) {
	if got := maxByte(); got != 255 {
		t.Errorf("maxByte() = %d, want 255", got)
	}
}
