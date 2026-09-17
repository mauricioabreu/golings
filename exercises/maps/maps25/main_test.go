// maps25
// Make the tests pass!

// I AM NOT DONE
//
// applyDiscount cuts the price of every item in the map by 10%.
// The prices do not change.
// The value in a range over a map is a copy, and it has to be written back.
package main_test

import "testing"

type Product struct {
	Name  string
	Price int
}

func applyDiscount(m map[string]Product) {
	for k, p := range m {
		p.Price = p.Price * 9 / 10
		_ = k
	}
}

func TestApplyDiscount(t *testing.T) {
	m := map[string]Product{"a": {"a", 100}, "b": {"b", 50}}
	applyDiscount(m)
	if m["a"].Price != 90 || m["b"].Price != 45 {
		t.Errorf("prices = %v", m)
	}
}
