// algorithms44
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: hash table and doubly linked list. Implement a cache with the
// given capacity: Get returns the value (or -1) and marks the key as most
// recently used, Put inserts or updates a value, evicting the least
// recently used key.
// Expected asymptotics: O(1) per operation, O(capacity) space.
package main_test

import "testing"

type node struct {
	key, val   int
	prev, next *node
}

type LRUCache struct {
	cap        int
	items      map[int]*node
	head, tail *node
}

func NewLRU(capacity int) *LRUCache {
	c := &LRUCache{cap: capacity, items: map[int]*node{}, head: &node{}, tail: &node{}}
	c.head.next, c.tail.prev = c.tail, c.head
	return c
}

func (c *LRUCache) unlink(n *node) { n.prev.next, n.next.prev = n.next, n.prev }

func (c *LRUCache) pushFront(n *node) {
	n.prev, n.next = c.head, c.head.next
	c.head.next.prev = n
	c.head.next = n
}

func (c *LRUCache) Get(key int) int {
	return -1
}

func (c *LRUCache) Put(key, val int) {
}

func TestLRU(t *testing.T) {
	c := NewLRU(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if c.Get(1) != 1 {
		t.Errorf("Get(1) should be 1")
	}
	c.Put(3, 3)
	if c.Get(2) != -1 {
		t.Errorf("2 should be evicted")
	}
	c.Put(4, 4)
	if c.Get(1) != -1 || c.Get(3) != 3 || c.Get(4) != 4 {
		t.Errorf("unexpected cache state")
	}
	c.Put(3, 30)
	c.Put(5, 5)
	if c.Get(3) != 30 || c.Get(4) != -1 {
		t.Errorf("update should refresh key 3")
	}
}
