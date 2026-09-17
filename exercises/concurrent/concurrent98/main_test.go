// concurrent98
// Make the tests pass!

// I AM NOT DONE
//
// Broadcaster sends a message to all subscribers over their channels.
// Practices a slice of channels guarded by a mutex.
package main_test

import (
	"sync"
	"testing"
)

type Broadcaster struct {
	mu   sync.Mutex
	subs []chan string
}

func (b *Broadcaster) Subscribe() <-chan string {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan string, 4)
	b.subs = []chan string{ch}
	return ch
}

func (b *Broadcaster) Publish(msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subs[0] <- msg
}

func TestBroadcaster(t *testing.T) {
	var b Broadcaster
	a := b.Subscribe()
	c := b.Subscribe()
	b.Publish("hi")
	if <-a != "hi" || <-c != "hi" {
		t.Errorf("not all subscribers got the message")
	}
}
