// concurrent35
// Make the tests pass!

// I AM NOT DONE
//
// fastest asks several mirrors and returns the first answer;
// the other goroutines must not hang forever.
// Practices a buffer large enough for every sender.
package main_test

import (
	"sync"
	"testing"
	"time"
)

func fastest(mirrors []func() string, wg *sync.WaitGroup) string {
	ch := make(chan string)
	for _, m := range mirrors {
		wg.Add(1)
		go func(f func() string) {
			ch <- f()
		}(m)
	}
	return <-ch
}

func TestFastest(t *testing.T) {
	var wg sync.WaitGroup
	slow := func() string { time.Sleep(30 * time.Millisecond); return "slow" }
	quick := func() string { return "quick" }
	if got := fastest([]func() string{slow, quick, slow}, &wg); got != "quick" {
		t.Errorf("fastest = %q", got)
	}
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("slow mirrors are blocked forever")
	}
}
