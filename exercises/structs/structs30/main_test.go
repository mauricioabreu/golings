// structs30
// Make the tests pass!

// I AM NOT DONE
//
// Stopwatch adds up the Start/Stop intervals, with the time passed in explicitly.
// Practices state in a struct plus time.Duration.
package main_test

import (
	"testing"
	"time"
)

type Stopwatch struct {
	started time.Time
	running bool
	total   time.Duration
}

func (s *Stopwatch) Start(now time.Time) {
	s.started = now
	s.running = true
}

func (s *Stopwatch) Stop(now time.Time) {
	s.total = now.Sub(s.started)
}

func TestStopwatch(t *testing.T) {
	base := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	var s Stopwatch
	s.Start(base)
	s.Stop(base.Add(2 * time.Second))
	s.Stop(base.Add(10 * time.Second))
	s.Start(base.Add(20 * time.Second))
	s.Stop(base.Add(23 * time.Second))
	if s.total != 5*time.Second {
		t.Errorf("total = %v, want 5s", s.total)
	}
}
