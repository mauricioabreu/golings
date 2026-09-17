// anonymous_functions60
// Make the tests pass!

// I AM NOT DONE
//
// byPrice sorts the items by price, keeping the order of equal ones.
// Practices a literal in sort.SliceStable.
package main_test

import (
	"sort"
	"testing"
)

type Item struct {
	Name  string
	Price int
}

func byPrice(items []Item) {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Price > items[j].Price
	})
}

func TestByPrice(t *testing.T) {
	items := []Item{{"b", 5}, {"a", 1}, {"c", 5}}
	byPrice(items)
	if items[0].Name != "a" || items[1].Name != "b" || items[2].Name != "c" {
		t.Errorf("byPrice = %v", items)
	}
}
