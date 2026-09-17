// variables94
// Make the tests pass!

// I AM NOT DONE
//
// This function must return the magnitude of the complex number 3+4i and its real part.
// Practices the complex128 type and the functions real, imag and cmplx.Abs.
package main_test

import (
	"math/cmplx"
	"testing"
)

func describe() (float64, float64) {
	z := complex(4, 4)
	return cmplx.Abs(z), imag(z)
}

func TestDescribe(t *testing.T) {
	abs, re := describe()
	if abs != 5 {
		t.Errorf("abs = %v, want 5", abs)
	}
	if re != 3 {
		t.Errorf("real part = %v, want 3", re)
	}
}
