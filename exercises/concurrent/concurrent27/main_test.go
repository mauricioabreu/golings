// concurrent27
// Make the tests pass!

// I AM NOT DONE
//
// handoff sends a task to a handler and checks right away that it was handled.
// With a buffered channel the send does not wait for the handler.
// An unbuffered channel is a rendezvous; a buffered one is not.
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Processor struct {
	mu   sync.Mutex
	done []int
	in   chan int
	ack  chan struct{}
}

func NewProcessor() *Processor {
	p := &Processor{in: make(chan int, 1), ack: make(chan struct{})}
	go func() {
		for v := range p.in {
			time.Sleep(10 * time.Millisecond)
			p.mu.Lock()
			p.done = append(p.done, v)
			p.mu.Unlock()
			p.ack <- struct{}{}
		}
	}()
	return p
}

func (p *Processor) Handoff(v int) {
	p.in <- v
	go func() { <-p.ack }()
}

func TestHandoff(t *testing.T) {
	p := NewProcessor()
	p.Handoff(7)
	p.mu.Lock()
	n := len(p.done)
	p.mu.Unlock()
	if n != 1 {
		t.Errorf("task not processed when Handoff returned (done = %d)", n)
	}
}
