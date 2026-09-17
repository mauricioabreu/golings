// maps10
// Make the tests pass!

// I AM NOT DONE
//
// The prices go into a map under a key built with Sprintf. When reading, the key
// is formatted differently and the price is not found.
// Keys have to be built by one and the same function.
package main_test

import (
	"fmt"
	"testing"
)

func priceKey(item string, size float64) string {
	return fmt.Sprintf("%s/%.1f", item, size)
}

type Prices map[string]int

func (p Prices) Set(item string, size float64, price int) {
	p[priceKey(item, size)] = price
}

func (p Prices) Get(item string, size float64) int {
	return p[fmt.Sprintf("%s/%v", item, size)]
}

func TestPrices(t *testing.T) {
	p := Prices{}
	p.Set("milk", 1, 90)
	p.Set("milk", 0.5, 50)
	if p.Get("milk", 1) != 90 || p.Get("milk", 0.5) != 50 {
		t.Errorf("prices = %v", p)
	}
}
