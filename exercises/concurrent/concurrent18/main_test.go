// concurrent18
// Make the tests pass!

// I AM NOT DONE
//
// SafeCounter keeps counters in a map under an RWMutex: Inc writes and Value reads.
// Practices Lock for writing and RLock for reading.
package main_test

import (
	"sync"
	"testing"
)

type SafeCounter struct {
	mu sync.RWMutex
	m  map[string]int
}

func (c *SafeCounter) Inc(k string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.m[k]++
}

func (c *SafeCounter) Value(k string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.m[k]
}

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{m: map[string]int{}}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("go")
			c.Value("go")
		}()
	}
	wg.Wait()
	if c.Value("go") != 100 {
		t.Errorf("value = %d", c.Value("go"))
	}
}
