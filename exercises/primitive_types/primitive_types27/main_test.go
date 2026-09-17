// primitive_types27
// Make the tests pass!

// I AM NOT DONE
//
// price must format a number with two decimal places.
// Practices the precision of the %f verb.
package main_test

import (
	"fmt"
	"testing"
)

func price(p float64) string {
	return fmt.Sprintf("%.1f", p)
}

func TestPrice(t *testing.T) {
	cases := map[float64]string{3.14159: "3.14", 2: "2.00", 0.005: "0.01"}
	for in, want := range cases {
		if got := price(in); got != want {
			t.Errorf("price(%v) = %s, want %s", in, got, want)
		}
	}
}
