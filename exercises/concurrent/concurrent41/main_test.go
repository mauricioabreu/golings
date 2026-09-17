// concurrent41
// Make the tests pass!

// I AM NOT DONE
//
// Group runs a function once for concurrent requests sharing a key,
// the way singleflight does; the rest wait and get the same result.
// Practices a map of waiters under a mutex plus a WaitGroup per request.
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type call struct {
	wg  sync.WaitGroup
	val string
}

type Group struct {
	mu    sync.Mutex
	calls map[string]*call
}

func (g *Group) Do(key string, fn func() string) string {
	g.mu.Lock()
	if g.calls == nil {
		g.calls = map[string]*call{}
	}
	g.mu.Unlock()
	return fn()
}

func TestGroup(t *testing.T) {
	var g Group
	var runs int32
	release := make(chan struct{})
	fn := func() string {
		atomic.AddInt32(&runs, 1)
		<-release
		return "data"
	}
	var wg sync.WaitGroup
	results := make([]string, 5)
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = g.Do("k", fn)
		}(i)
	}
	time.Sleep(50 * time.Millisecond)
	close(release)
	wg.Wait()
	if atomic.LoadInt32(&runs) != 1 {
		t.Errorf("fn ran %d times", runs)
	}
	for _, r := range results {
		if r != "data" {
			t.Errorf("result = %q", r)
		}
	}
}
