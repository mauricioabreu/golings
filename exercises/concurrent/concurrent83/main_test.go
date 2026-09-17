// concurrent83
// Make the tests pass!

// I AM NOT DONE
//
// update releases the lock in two places on early exit, and the program
// crashes with "unlock of unlocked mutex".
// Practices one Lock/Unlock pair per execution path.
package main_test

import (
	"sync"
	"testing"
)

type Cache struct {
	mu   sync.Mutex
	data map[string]string
}

func (c *Cache) Update(k, v string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v == "" {
		c.mu.Unlock()
		return false
	}
	c.data[k] = v
	return true
}

func TestUpdate(t *testing.T) {
	c := &Cache{data: map[string]string{}}
	if !c.Update("a", "1") || c.Update("b", "") || !c.Update("c", "3") {
		t.Errorf("Update results are wrong")
	}
}
