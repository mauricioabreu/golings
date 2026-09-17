// primitive_types89
// Make the tests pass!

// I AM NOT DONE
//
// toSigned must convert a uint64 to an int64 and return an error
// when the value does not fit.
// A conversion between a signed and an unsigned type does not check the range.
package main_test

import (
	"errors"
	"math"
	"testing"
)

var ErrRange = errors.New("out of range")

func toSigned(v uint64) (int64, error) {
	if int64(v) > math.MaxInt64 {
		return 0, ErrRange
	}
	return int64(v), nil
}

func TestToSigned(t *testing.T) {
	if got, err := toSigned(42); err != nil || got != 42 {
		t.Errorf("toSigned(42) = %d, %v", got, err)
	}
	if got, err := toSigned(math.MaxInt64); err != nil || got != math.MaxInt64 {
		t.Errorf("toSigned(max) = %d, %v", got, err)
	}
	if _, err := toSigned(math.MaxUint64); !errors.Is(err, ErrRange) {
		t.Errorf("toSigned(MaxUint64) error = %v, want ErrRange", err)
	}
}
