// generics6
// Make the tests pass!

// I AM NOT DONE
//
// Cache[K, V] holds no more than cap entries and drops the oldest key
// when it is full (FIFO).
// Practices a generic type with a map and a slice for the order.
package main_test

import "testing"

type Cache[K comparable, V any] struct {
	cap   int
	data  map[K]V
	order []K
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{cap: capacity, data: map[K]V{}}
}

func (c *Cache[K, V]) Put(k K, v V) {
	c.data[k] = v
	c.order = append(c.order, k)
}

func (c *Cache[K, V]) Get(k K) (V, bool) {
	v, ok := c.data[k]
	return v, ok
}

func TestCache(t *testing.T) {
	c := NewCache[string, int](2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("a", 10)
	c.Put("c", 3)
	if _, ok := c.Get("a"); ok {
		t.Errorf("a should be evicted")
	}
	if v, ok := c.Get("b"); !ok || v != 2 {
		t.Errorf("b = %d, %v", v, ok)
	}
	if len(c.data) != 2 {
		t.Errorf("size = %d", len(c.data))
	}
}
