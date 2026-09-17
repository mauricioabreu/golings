// slices67
// Make the tests pass!

// I AM NOT DONE
//
// restore copies a saved state into the working buffer and has to report
// an error when the buffer is too small. Right now the data is quietly truncated.
// copy does not report a truncation, so the result has to be checked.
package main_test

import (
	"errors"
	"testing"
)

func restore(dst, saved []int) error {
	copy(dst, saved)
	return nil
}

func TestRestore(t *testing.T) {
	_ = errors.New
	dst := make([]int, 3)
	if err := restore(dst, []int{1, 2, 3}); err != nil || dst[2] != 3 {
		t.Errorf("restore into big buffer: %v, %v", err, dst)
	}
	if err := restore(make([]int, 2), []int{1, 2, 3}); err == nil {
		t.Errorf("restore into small buffer should fail")
	}
}
