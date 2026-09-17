// if37
// Make the tests pass!

// I AM NOT DONE
//
// shippingCost: up to 1 kg is 200, up to 5 kg is 400, heavier is 400 plus 100 for every
// whole kilogram over 5. Express delivery doubles the price.
// Practices combining a branch with a modifier.
package main_test

import "testing"

func shippingCost(kg int, express bool) int {
	var cost int
	if kg <= 1 {
		cost = 200
	} else if kg <= 5 {
		cost = 400
	} else {
		cost = 100 * (kg - 5)
	}
	if express {
		cost += 2
	}
	return cost
}

func TestShippingCost(t *testing.T) {
	cases := []struct {
		kg      int
		express bool
		want    int
	}{{1, false, 200}, {3, false, 400}, {8, false, 700}, {1, true, 400}, {8, true, 1400}}
	for _, c := range cases {
		if got := shippingCost(c.kg, c.express); got != c.want {
			t.Errorf("shippingCost(%d, %v) = %d, want %d", c.kg, c.express, got, c.want)
		}
	}
}
