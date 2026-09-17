// primitive_types77
// Make the tests pass!

// I AM NOT DONE
//
// total adds up prices in whole units with fractions and returns the sum in cents.
// The sum of ten prices of 0.10 is not 100 cents.
// Practices float64 error building up in money calculations.
package main_test

import (
	"math"
	"testing"
)

func total(prices []float64) int64 {
	var sum float64
	for _, p := range prices {
		sum += p
	}
	return int64(sum * 100)
}

func TestTotal(t *testing.T) {
	_ = math.Round
	ten := make([]float64, 10)
	for i := range ten {
		ten[i] = 0.10
	}
	if got := total(ten); got != 100 {
		t.Errorf("total(10 x 0.10) = %d, want 100", got)
	}
	if got := total([]float64{19.99, 0.01, 4.35}); got != 2435 {
		t.Errorf("total = %d, want 2435", got)
	}
}
