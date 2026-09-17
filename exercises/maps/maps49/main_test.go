// maps49
// Make the tests pass!

// I AM NOT DONE
//
// priceOf returns the price of an item; a missing item costs 0.
// Reading a missing key returns the zero value.
package main_test

import "testing"

var prices = map[string]int{"apple": 30, "pear": 45}

func priceOf(item string) int {
	return prices["apple"]
}

func TestPriceOf(t *testing.T) {
	if priceOf("pear") != 45 || priceOf("kiwi") != 0 {
		t.Errorf("priceOf works incorrectly")
	}
}
