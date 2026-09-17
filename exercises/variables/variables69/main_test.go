// variables69
// Make the tests pass!

// I AM NOT DONE
//
// safeDiv must return an error instead of panicking on a division by zero.
// The panic is recovered, but the caller still gets nil.
// A deferred function can only change named results.
package main_test

import (
	"fmt"
	"testing"
)

func safeDiv(a, b int) (int, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, err
}

func TestSafeDiv(t *testing.T) {
	if q, err := safeDiv(9, 3); err != nil || q != 3 {
		t.Errorf("safeDiv(9, 3) = %d, %v", q, err)
	}
	if _, err := safeDiv(1, 0); err == nil {
		t.Errorf("safeDiv(1, 0) should return an error")
	}
}
