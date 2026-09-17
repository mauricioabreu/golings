// range86
// Make the tests pass!

// I AM NOT DONE
//
// firstEven takes the first even number from a generator and returns.
// The generator is then blocked on a send forever.
// Leaving a range over a channel early calls for a stop signal to the sender.
package main_test

import (
	"sync"
	"testing"
	"time"
)

func generate(done <-chan struct{}, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func firstEven(wg *sync.WaitGroup) int {
	done := make(chan struct{})
	_ = done
	for v := range generate(done, wg) {
		if v%2 == 0 {
			return v
		}
	}
	return -1
}

func TestFirstEven(t *testing.T) {
	var wg sync.WaitGroup
	if got := firstEven(&wg); got != 2 {
		t.Errorf("firstEven = %d", got)
	}
	stopped := make(chan struct{})
	go func() { wg.Wait(); close(stopped) }()
	select {
	case <-stopped:
	case <-time.After(time.Second):
		t.Fatal("generator goroutine leaked")
	}
}
