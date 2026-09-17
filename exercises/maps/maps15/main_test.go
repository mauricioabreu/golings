// maps15
// Make the tests pass!

// I AM NOT DONE
//
// Cache.Put stores a value, but the method has a value receiver
// and the map is built inside the method. The data is lost.
// Assigning to a map field in a method with a value receiver is invisible outside.
package main_test

import "testing"

type Cache struct {
	items map[string]string
}

func (c Cache) Put(k, v string) {
	if c.items == nil {
		c.items = map[string]string{}
	}
	c.items[k] = v
}

func (c Cache) Get(k string) string { return c.items[k] }

func TestCache(t *testing.T) {
	var c Cache
	c.Put("a", "1")
	c.Put("b", "2")
	if c.Get("a") != "1" || c.Get("b") != "2" {
		t.Errorf("cache lost values: %v", c.items)
	}
}
