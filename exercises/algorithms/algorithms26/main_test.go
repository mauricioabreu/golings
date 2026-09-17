// algorithms26
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: hash tables and lists by frequency. A cache of the given
// capacity evicts the key with the fewest accesses; on a tie — the least
// recently used. Get and Put increase the key's frequency.
// Expected asymptotics: O(1) per operation, O(capacity) space.
package main_test

import (
	"container/list"
	"testing"
)

type entry struct {
	key, val, freq int
	elem           *list.Element
}

type LFUCache struct {
	cap     int
	minFreq int
	items   map[int]*entry
	freqs   map[int]*list.List
}

func NewLFU(capacity int) *LFUCache {
	return &LFUCache{cap: capacity, items: map[int]*entry{}, freqs: map[int]*list.List{}}
}

func (c *LFUCache) touch(e *entry) {
}

func (c *LFUCache) Get(key int) int {
	e, ok := c.items[key]
	if !ok {
		return -1
	}
	c.touch(e)
	return e.val
}

func (c *LFUCache) Put(key, val int) {
	if c.cap == 0 {
		return
	}
	if e, ok := c.items[key]; ok {
		e.val = val
		c.touch(e)
		return
	}
	c.items[key] = &entry{key: key, val: val}
}

func TestLFU(t *testing.T) {
	c := NewLFU(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if c.Get(1) != 1 {
		t.Errorf("Get(1) = %d", c.Get(1))
	}
	c.Put(3, 3)
	if c.Get(2) != -1 || c.Get(3) != 3 {
		t.Errorf("key 2 should be evicted")
	}
	c.Put(4, 4)
	if c.Get(1) != -1 || c.Get(3) != 3 || c.Get(4) != 4 {
		t.Errorf("key 1 should be evicted (tie broken by recency)")
	}
	z := NewLFU(0)
	z.Put(1, 1)
	if z.Get(1) != -1 {
		t.Errorf("zero-capacity cache stores values")
	}
}
