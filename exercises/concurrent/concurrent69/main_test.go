// concurrent69
// Make the tests pass!

// I AM NOT DONE
//
// Instance creates a singleton with "double-checked locking": the first
// check happens without a lock and causes a race.
// Practices that reading a shared field also requires synchronization.
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Holder struct {
	mu   sync.Mutex
	once sync.Once
	inst *string
}

var created int32

func (h *Holder) Instance() *string {
	if h.inst == nil {
		h.mu.Lock()
		if h.inst == nil {
			s := "service"
			atomic.AddInt32(&created, 1)
			h.inst = &s
		}
		h.mu.Unlock()
	}
	return h.inst
}

func TestInstance(t *testing.T) {
	var h Holder
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if *h.Instance() != "service" {
				t.Errorf("wrong instance")
			}
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&created) != 1 {
		t.Errorf("created = %d", created)
	}
}
