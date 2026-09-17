// slices64
// Make the tests pass!

// I AM NOT DONE
//
// Queue implements Enqueue and Dequeue (FIFO).
// Practices the slice [1:] for dropping the first element.
package main_test

import "testing"

type Queue struct{ items []string }

func (q *Queue) Enqueue(v string) { q.items = append(q.items, v) }

func (q *Queue) Dequeue() (string, bool) {
	v := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return v, true
}

func TestQueue(t *testing.T) {
	var q Queue
	q.Enqueue("a")
	q.Enqueue("b")
	if v, _ := q.Dequeue(); v != "a" {
		t.Errorf("Dequeue = %q, want a", v)
	}
	if v, _ := q.Dequeue(); v != "b" {
		t.Errorf("Dequeue = %q, want b", v)
	}
	if _, ok := q.Dequeue(); ok {
		t.Errorf("Dequeue on empty queue should fail")
	}
}
