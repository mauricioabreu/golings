// concurrent13
// Make the tests pass!

// I AM NOT DONE
//
// The counter is raised through atomic and read directly.
// The race detector reports a problem.
// Every access to an atomic variable has to go through atomic.
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Meter struct{ bytes int64 }

func (m *Meter) Add(n int64) { atomic.AddInt64(&m.bytes, n) }

func (m *Meter) Total() int64 {
	return m.bytes
}

func TestMeter(t *testing.T) {
	var m Meter
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() { defer wg.Done(); m.Add(10) }()
		go func() { defer wg.Done(); _ = m.Total() }()
	}
	wg.Wait()
	if m.Total() != 500 {
		t.Errorf("Total = %d", m.Total())
	}
}
