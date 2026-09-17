// structs75
// Make the tests pass!

// I AM NOT DONE
//
// A copy of a struct is independent of the original.
// withDiscount returns a discounted item without changing the original.
// Assigning a struct copies it.
package main_test

import "testing"

type Item struct {
	Name  string
	Price int
}

func withDiscount(it *Item) Item {
	c := *it
	it.Price /= 2
	return c
}

func TestWithDiscount(t *testing.T) {
	it := Item{"tea", 100}
	d := withDiscount(&it)
	if d.Price != 50 || it.Price != 100 {
		t.Errorf("discounted = %d, original = %d", d.Price, it.Price)
	}
}
