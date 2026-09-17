// concurrent8
// Make the tests pass!

// I AM NOT DONE
//
// The Add method has a value receiver: every call locks a copy
// of the mutex and changes a copy of the data.
// Structs with a mutex use a pointer receiver.
package main_test

import (
	"sync"
	"testing"
)

type Tally struct {
	mu    sync.Mutex
	total *int
}

func (t Tally) Add(n int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	*t.total += n
}

func TestTally(t *testing.T) {
	total := 0
	tally := &Tally{total: &total}
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tally.Add(1)
		}()
	}
	wg.Wait()
	if total != 200 {
		t.Errorf("total = %d", total)
	}
}
