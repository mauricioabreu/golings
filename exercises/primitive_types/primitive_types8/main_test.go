// primitive_types8
// Make the tests pass!

// I AM NOT DONE
//
// product multiplies the factors and must return an error when the result
// stops being a finite number.
// A float64 overflow gives +Inf without a panic.
package main_test

import (
	"errors"
	"math"
	"testing"
)

func product(xs []float64) (float64, error) {
	p := 1.0
	for _, x := range xs {
		p *= x
		if p > math.MaxFloat64 {
			return 0, errors.New("overflow")
		}
	}
	return p, nil
}

func TestProduct(t *testing.T) {
	if v, err := product([]float64{2, 3, 4}); err != nil || v != 24 {
		t.Errorf("product(2,3,4) = %v, %v", v, err)
	}
	if _, err := product([]float64{1e200, 1e200}); err == nil {
		t.Errorf("product(1e200, 1e200) should overflow")
	}
	if _, err := product([]float64{-1e200, 1e200}); err == nil {
		t.Errorf("product(-1e200, 1e200) should overflow")
	}
}
