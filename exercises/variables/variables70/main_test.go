// variables70
// Make the tests pass!

// I AM NOT DONE
//
// formatPrice must return the price with exactly two decimal places.
// Practices strconv.FormatFloat and its parameters.
package main_test

import (
	"strconv"
	"testing"
)

func formatPrice(p float64) string {
	return strconv.FormatFloat(p, 'e', -1, 64)
}

func TestFormatPrice(t *testing.T) {
	cases := map[float64]string{9.5: "9.50", 10: "10.00", 0.125: "0.12", 1234.567: "1234.57"}
	for in, want := range cases {
		if got := formatPrice(in); got != want {
			t.Errorf("formatPrice(%v) = %q, want %q", in, got, want)
		}
	}
}
