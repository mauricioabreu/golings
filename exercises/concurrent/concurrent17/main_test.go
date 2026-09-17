// concurrent17
// Make the tests pass!

// I AM NOT DONE
//
// init calls Get, which enters the same once.Do again, which is a deadlock.
// Calling Do from inside the f of the same Once blocks forever.
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Lazy struct {
	once sync.Once
	vals []string
}

func (l *Lazy) Get() []string {
	l.once.Do(func() {
		l.vals = []string{"a", "b"}
		if len(l.Get()) == 0 {
			l.vals = []string{"default"}
		}
	})
	return l.vals
}

func TestLazy(t *testing.T) {
	var l Lazy
	res := make(chan []string, 1)
	go func() { res <- l.Get() }()
	select {
	case v := <-res:
		if len(v) != 2 {
			t.Errorf("Get = %v", v)
		}
	case <-time.After(time.Second):
		t.Fatal("Get deadlocked inside once.Do")
	}
}
