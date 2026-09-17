// concurrent59
// Make the tests pass!

// I AM NOT DONE
//
// Snapshot reads under RLock and inside calls a method that takes
// RLock again.
// If a writer is waiting on Lock at that moment, the second RLock
// blocks — deadlock.
// Practices that RWMutex does not allow recursive read locks.
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Store struct {
	mu    sync.RWMutex
	items []string
}

func (s *Store) count() int { return len(s.items) }

func (s *Store) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.count()
}

func (s *Store) Summary(pause chan struct{}) (int, string) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	close(pause)
	time.Sleep(50 * time.Millisecond)
	return s.Count(), s.items[0]
}

func (s *Store) Add(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, v)
}

func TestSummary(t *testing.T) {
	s := &Store{items: []string{"x"}}
	pause := make(chan struct{})
	res := make(chan int, 1)
	go func() {
		n, _ := s.Summary(pause)
		res <- n
	}()
	<-pause
	go s.Add("y")
	select {
	case n := <-res:
		if n != 1 {
			t.Errorf("Summary count = %d", n)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("deadlock: recursive RLock with a pending writer")
	}
}
