// concurrent9
// Make the tests pass!

package main_test

import (
	"sync"
	"testing"
	"time"
)

// merge sends everything that arrives on any of the channels to the channel it
// returns, and closes that channel once all of the inputs are closed.
func merge(inputs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range in {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func receive(t *testing.T, ch <-chan int) (int, bool) {
	t.Helper()
	select {
	case v, ok := <-ch:
		return v, ok
	case <-time.After(time.Second):
		t.Fatal("waited a second for the merged channel")
		return 0, false
	}
}

func TestMergeDoesNotWaitForTheFirstInput(t *testing.T) {
	slow := make(chan int)
	fast := make(chan int)
	release := make(chan struct{})

	go func() {
		<-release
		slow <- 1
		close(slow)
	}()
	go func() {
		fast <- 2
		close(fast)
	}()

	merged := merge(slow, fast)

	if v, _ := receive(t, merged); v != 2 {
		t.Fatalf("first value = %d, want 2: merge is waiting for the first input", v)
	}
	close(release)
	if v, _ := receive(t, merged); v != 1 {
		t.Errorf("second value = %d, want 1", v)
	}
}

func TestMergeClosesWhenAllInputsAreClosed(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	go func() { a <- 1; close(a) }()
	go func() { b <- 2; close(b) }()

	merged := merge(a, b)

	sum := 0
	for range 2 {
		v, _ := receive(t, merged)
		sum += v
	}
	if sum != 3 {
		t.Errorf("sum of the values = %d, want 3", sum)
	}
	if _, ok := receive(t, merged); ok {
		t.Error("merged should be closed after the last value")
	}
}
