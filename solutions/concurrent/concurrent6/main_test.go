// concurrent6
// Make the tests pass!

package main_test

import (
	"testing"
	"time"
)

// firstResult returns what arrives on results, or "timeout" if nothing arrives
// within the limit.
func firstResult(results <-chan string, limit time.Duration) string {
	select {
	case result := <-results:
		return result
	case <-time.After(limit):
		return "timeout"
	}
}

// slowWorker sends its result after the delay. The channel has room for it, so
// the worker never gets stuck if nobody is listening anymore.
func slowWorker(delay time.Duration, result string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(delay)
		ch <- result
	}()
	return ch
}

func TestFirstResultArrivesInTime(t *testing.T) {
	got := firstResult(slowWorker(5*time.Millisecond, "done"), time.Second)

	if got != "done" {
		t.Errorf("got %q, want %q", got, "done")
	}
}

func TestFirstResultGivesUp(t *testing.T) {
	started := time.Now()

	got := firstResult(slowWorker(500*time.Millisecond, "late"), 50*time.Millisecond)

	if got != "timeout" {
		t.Errorf("got %q, want %q", got, "timeout")
	}
	if elapsed := time.Since(started); elapsed > 400*time.Millisecond {
		t.Errorf("gave up after %v, want about 50ms", elapsed)
	}
}
