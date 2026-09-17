// concurrent60
// Make the tests pass!

// I AM NOT DONE
//
// BlockingQueue.Take waits until an element appears; Put adds one and
// wakes up waiters.
// Practices sync.Cond.
package main_test

import (
	"sync"
	"testing"
	"time"
)

type BlockingQueue struct {
	mu    sync.Mutex
	cond  *sync.Cond
	items []int
}

func NewQueue() *BlockingQueue {
	q := &BlockingQueue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *BlockingQueue) Put(v int) {
	q.mu.Lock()
	q.items = append(q.items, v)
	q.mu.Unlock()
}

func (q *BlockingQueue) Take() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait()
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v
}

func TestBlockingQueue(t *testing.T) {
	q := NewQueue()
	got := make(chan int, 1)
	go func() { got <- q.Take() }()
	time.Sleep(20 * time.Millisecond)
	q.Put(7)
	select {
	case v := <-got:
		if v != 7 {
			t.Errorf("Take = %d", v)
		}
	case <-time.After(time.Second):
		t.Fatal("Take was never woken up")
	}
}
