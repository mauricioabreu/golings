// maps42
// Make the tests pass!

// I AM NOT DONE
//
// The price cache is indexed by a string glued together from a category and an article number.
// Different pairs give the same key and overwrite each other.
// Keys built by concatenating strings are ambiguous.
package main_test

import "testing"

type PriceCache map[string]int

func (c PriceCache) Set(cat, sku string, p int) { c[cat+sku] = p }
func (c PriceCache) Get(cat, sku string) int    { return c[cat+sku] }

func TestPriceCache(t *testing.T) {
	c := PriceCache{}
	c.Set("tv", "100", 500)
	c.Set("tv1", "00", 900)
	if c.Get("tv", "100") != 500 || c.Get("tv1", "00") != 900 {
		t.Errorf("prices collided: %v", c)
	}
}
