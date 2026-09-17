// concurrent66
// Make the tests pass!

// I AM NOT DONE
//
// Publisher sends an event to a subscriber while holding a mutex; the
// subscriber needs the same mutex to process it — deadlock.
// Practices not performing blocking operations under a lock.
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Publisher struct {
	mu     sync.Mutex
	events chan string
	sent   int
}

func (p *Publisher) Send(e string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.sent++
	p.events <- e
}

func (p *Publisher) Sent() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.sent
}

func TestPublisher(t *testing.T) {
	p := &Publisher{events: make(chan string)}
	got := make(chan int, 1)
	go func() {
		<-p.events
		time.Sleep(20 * time.Millisecond)
		p.Sent()
		<-p.events
		got <- p.Sent()
	}()
	go func() {
		p.Send("x")
		p.Send("y")
	}()
	select {
	case n := <-got:
		if n != 2 {
			t.Errorf("Sent = %d", n)
		}
	case <-time.After(time.Second):
		t.Fatal("deadlock: sending while holding the mutex")
	}
}
