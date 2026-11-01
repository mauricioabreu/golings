// concurrent2
// Make the tests pass!

package main_test

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	counter := updateCounter()
	if counter != 100 {
		t.Errorf("Counter should be 100, but got %d", counter)
	}
}

func updateCounter() int {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++ // Many goroutines trying to update the counter? We need some protection here!
			mu.Unlock()
		}()
	}
	wg.Wait()
	return counter
}
