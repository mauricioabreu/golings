// arrays29
// Make the tests pass!

// I AM NOT DONE
//
// discountAll cuts the price of every item by 10%.
// The prices in the array do not change.
// The value variable of a range is a copy of the element.
package main_test

import "testing"

type Item struct {
	Name  string
	Price int
}

func discountAll(items *[3]Item) {
	for _, it := range items {
		it.Price = it.Price * 9 / 10
	}
}

func TestDiscountAll(t *testing.T) {
	items := [3]Item{{"a", 100}, {"b", 250}, {"c", 10}}
	discountAll(&items)
	if items[0].Price != 90 || items[1].Price != 225 || items[2].Price != 9 {
		t.Errorf("prices = %d %d %d", items[0].Price, items[1].Price, items[2].Price)
	}
}
