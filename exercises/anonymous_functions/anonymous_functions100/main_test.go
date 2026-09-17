// anonymous_functions100
// Make the tests pass!

// I AM NOT DONE
//
// pricing returns a price literal for a strategy name;
// an unknown name gives the strategy that changes nothing.
// Practices returning different literals from a function.
package main_test

import "testing"

func pricing(name string) func(price, qty int) int {
	switch name {
	case "sale":
		return func(price, qty int) int { return price * qty * 20 / 100 }
	case "bulk":
		return func(price, qty int) int { return price*qty - 10 }
	}
	return func(price, qty int) int { return price * qty }
}

func TestPricing(t *testing.T) {
	if pricing("sale")(100, 2) != 160 || pricing("bulk")(10, 25) != 230 || pricing("none")(3, 3) != 9 {
		t.Errorf("pricing: %d %d %d", pricing("sale")(100, 2), pricing("bulk")(10, 25), pricing("none")(3, 3))
	}
}
