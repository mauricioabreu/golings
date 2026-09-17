// concurrent101
// Make the tests pass!

// I AM NOT DONE
//
// Stats counts HTTP codes from many goroutines; Snapshot returns a map copy.
// Practices a mutex and returning a copy of the guarded data.
package main_test

import (
	"sync"
	"testing"
)

type Stats struct {
	mu    sync.Mutex
	codes map[int]int
}

func (s *Stats) Add(code int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.codes[code]++
}

func (s *Stats) Snapshot() map[int]int {
	return s.codes
}

func TestStats(t *testing.T) {
	s := &Stats{codes: map[int]int{}}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(200 + i%2*300)
			_ = len(s.Snapshot())
		}(i)
	}
	wg.Wait()
	snap := s.Snapshot()
	snap[200] = -1
	if s.Snapshot()[200] != 25 || s.Snapshot()[500] != 25 {
		t.Errorf("stats = %v", s.Snapshot())
	}
}
