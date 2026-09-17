// concurrent80
// Make the tests pass!

// I AM NOT DONE
//
// Take waits for an element using sync.Cond, but after waking it does not
// recheck the condition. On Broadcast, several consumers wake for one item.
// Practices calling Wait in a loop.
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Queue struct {
	mu    sync.Mutex
	cond  *sync.Cond
	items []int
}

func NewQueue() *Queue {
	q := &Queue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *Queue) Put(v int) {
	q.mu.Lock()
	q.items = append(q.items, v)
	q.mu.Unlock()
	q.cond.Broadcast()
}

func (q *Queue) Take() (v int, ok bool) {
	defer func() {
		if recover() != nil {
			ok = false
		}
	}()
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		q.cond.Wait()
	}
	v = q.items[0]
	q.items = q.items[1:]
	return v, true
}

func TestTakeLoop(t *testing.T) {
	q := NewQueue()
	results := make(chan bool, 3)
	for i := 0; i < 3; i++ {
		go func() {
			_, ok := q.Take()
			results <- ok
		}()
	}
	time.Sleep(30 * time.Millisecond)
	q.Put(1)
	time.Sleep(30 * time.Millisecond)
	q.Put(2)
	q.Put(3)
	for i := 0; i < 3; i++ {
		select {
		case ok := <-results:
			if !ok {
				t.Fatalf("a consumer woke up with an empty queue")
			}
		case <-time.After(time.Second):
			t.Fatal("consumers are stuck")
		}
	}
}
