// concurrent93
// Make the tests pass!

// I AM NOT DONE
//
// Stop may be called from several places; a second close panics.
// Practices closing a channel exactly once.
package main_test

import (
	"sync"
	"testing"
)

type Worker struct {
	once sync.Once
	quit chan struct{}
}

func (w *Worker) Stop() {
	close(w.quit)
}

func TestStopTwice(t *testing.T) {
	w := &Worker{quit: make(chan struct{})}
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Stop panicked: %v", r)
				}
			}()
			w.Stop()
		}()
	}
	wg.Wait()
	select {
	case <-w.quit:
	default:
		t.Errorf("quit channel is not closed")
	}
}
