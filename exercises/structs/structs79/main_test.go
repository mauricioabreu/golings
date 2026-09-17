// structs79
// Make the tests pass!

// I AM NOT DONE
//
// priceOf reads the price of an item from a map of structs.
// Practices reading a field of a map value.
package main_test

import "testing"

type Product struct {
	Price int
	Stock int
}

func priceOf(m map[string]Product, name string) int {
	return m[name].Stock
}

func TestPriceOf(t *testing.T) {
	m := map[string]Product{"pen": {Price: 30, Stock: 7}}
	if priceOf(m, "pen") != 30 || priceOf(m, "cup") != 0 {
		t.Errorf("priceOf works incorrectly")
	}
}
