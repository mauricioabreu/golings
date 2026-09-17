// concurrent78
// Make the tests pass!

// I AM NOT DONE
//
// Gate.Open must wake all waiters, but Signal wakes only one.
// Practices the difference between Signal and Broadcast.
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Gate struct {
	mu   sync.Mutex
	cond *sync.Cond
	open bool
}

func NewGate() *Gate {
	g := &Gate{}
	g.cond = sync.NewCond(&g.mu)
	return g
}

func (g *Gate) Wait() {
	g.mu.Lock()
	for !g.open {
		g.cond.Wait()
	}
	g.mu.Unlock()
}

func (g *Gate) Open() {
	g.mu.Lock()
	g.open = true
	g.mu.Unlock()
	g.cond.Signal()
}

func TestGate(t *testing.T) {
	g := NewGate()
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); g.Wait() }()
	}
	time.Sleep(30 * time.Millisecond)
	g.Open()
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("not all waiters were woken")
	}
}
