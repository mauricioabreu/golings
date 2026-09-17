// if59
// Make the tests pass!

// I AM NOT DONE
//
// Limiter allows no more than max requests in a window of window seconds.
// When the window runs out, the counter is reset.
// Practices conditions that change the state of a struct.
package main_test

import "testing"

type Limiter struct {
	max, window int
	start, used int
}

func (l *Limiter) Allow(now int) bool {
	if l.used > l.max {
		return false
	}
	l.used++
	return true
}

func TestLimiter(t *testing.T) {
	l := &Limiter{max: 2, window: 10}
	got := []bool{l.Allow(0), l.Allow(1), l.Allow(2), l.Allow(10), l.Allow(11), l.Allow(12)}
	want := []bool{true, true, false, true, true, false}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("call %d: Allow = %v, want %v", i, got[i], want[i])
		}
	}
}
