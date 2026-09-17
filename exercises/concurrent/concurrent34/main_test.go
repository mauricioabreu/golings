// concurrent34
// Make the tests pass!

// I AM NOT DONE
//
// getOrCreate holds an RLock and tries to take the write Lock, which is a deadlock.
// An RWMutex cannot upgrade a lock.
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Registry struct {
	mu sync.RWMutex
	m  map[string]int
}

func (r *Registry) GetOrCreate(k string) int {
	r.mu.RLock()
	v, ok := r.m[k]
	if !ok {
		r.mu.Lock()
		r.m[k] = len(r.m) + 1
		v = r.m[k]
		r.mu.Unlock()
	}
	r.mu.RUnlock()
	return v
}

func TestGetOrCreate(t *testing.T) {
	r := &Registry{m: map[string]int{}}
	res := make(chan int, 1)
	go func() { res <- r.GetOrCreate("a") }()
	select {
	case v := <-res:
		if v != 1 || r.GetOrCreate("a") != 1 {
			t.Errorf("GetOrCreate = %d", v)
		}
	case <-time.After(time.Second):
		t.Fatal("GetOrCreate deadlocked")
	}
}
