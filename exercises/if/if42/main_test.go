// if42
// Make the tests pass!

// I AM NOT DONE
//
// discountFor must return the discount from the table, or 0 when the code is missing.
// Practices an if with an init statement and comma-ok.
package main_test

import "testing"

var discounts = map[string]int{"SALE10": 10, "FREE": 100}

func discountFor(code string) int {
	if d, ok := discounts[code]; !ok {
		return d
	}
	return 0
}

func TestDiscountFor(t *testing.T) {
	cases := map[string]int{"SALE10": 10, "FREE": 100, "NOPE": 0}
	for in, want := range cases {
		if got := discountFor(in); got != want {
			t.Errorf("discountFor(%q) = %d, want %d", in, got, want)
		}
	}
}
