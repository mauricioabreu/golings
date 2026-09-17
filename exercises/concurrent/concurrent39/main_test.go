// concurrent39
// Make the tests pass!

// I AM NOT DONE
//
// reserve atomically lowers the number of seats while it is above zero.
// Practices atomic.CompareAndSwapInt32 in a loop.
package main_test

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func reserve(seats *int32) bool {
	for {
		cur := atomic.LoadInt32(seats)
		if cur <= 0 {
			return false
		}
		runtime.Gosched()
		atomic.StoreInt32(seats, cur-1)
		return true
	}
}

func TestReserve(t *testing.T) {
	seats := int32(50)
	var ok int32
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if reserve(&seats) {
				atomic.AddInt32(&ok, 1)
			}
		}()
	}
	wg.Wait()
	if ok != 50 || seats != 0 {
		t.Errorf("reserved %d, left %d", ok, seats)
	}
}
