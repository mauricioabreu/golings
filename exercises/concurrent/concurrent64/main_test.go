// concurrent64
// Make the tests pass!

// I AM NOT DONE
//
// Visits.Hit is called from many goroutines; the counter loses updates
// and the race detector reports an error.
// Practices protecting a shared variable with a mutex.
package main_test

import (
	"sync"
	"testing"
)

type Visits struct {
	mu sync.Mutex
	n  int
}

func (v *Visits) Hit() {
	v.n++
}

func TestVisits(t *testing.T) {
	var v Visits
	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v.Hit()
		}()
	}
	wg.Wait()
	if v.n != 500 {
		t.Errorf("n = %d, want 500", v.n)
	}
}
