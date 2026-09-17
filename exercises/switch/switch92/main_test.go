// switch92
// Make the tests pass!

// I AM NOT DONE
//
// shipping: orders of 5000 and up ship free; orders of 1000 and up cost 100;
// anything else 300. A fallthrough produces the wrong price.
// fallthrough moves to the next branch without checking its condition.
package main_test

import "testing"

func shipping(total int) int {
	cost := 300
	switch {
	case total >= 5000:
		cost = 0
		fallthrough
	case total >= 1000:
		cost = 100
	}
	return cost
}

func TestShipping(t *testing.T) {
	cases := map[int]int{6000: 0, 5000: 0, 1500: 100, 999: 300}
	for in, want := range cases {
		if got := shipping(in); got != want {
			t.Errorf("shipping(%d) = %d, want %d", in, got, want)
		}
	}
}
