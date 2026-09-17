// primitive_types68
// Make the tests pass!

// I AM NOT DONE
//
// roots solves the quadratic equation ax²+bx+c=0 and always returns two roots
// as complex128 values, a negative discriminant included.
// Practices complex128 and the math/cmplx package.
package main_test

import (
	"math/cmplx"
	"testing"
)

func roots(a, b, c float64) (complex128, complex128) {
	d := b*b - 4*a*c
	sq := complex(d, 0)
	return (complex(-b, 0) + sq) / complex(a, 0), (complex(-b, 0) - sq) / complex(a, 0)
}

func TestRoots(t *testing.T) {
	_ = cmplx.Sqrt
	x1, x2 := roots(1, -3, 2)
	if x1 != 2 || x2 != 1 {
		t.Errorf("roots(1,-3,2) = %v, %v", x1, x2)
	}
	y1, y2 := roots(1, 0, 4)
	if y1 != 2i || y2 != -2i {
		t.Errorf("roots(1,0,4) = %v, %v; want 2i, -2i", y1, y2)
	}
}
