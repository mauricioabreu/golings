// structs84
// Make the tests pass!

// I AM NOT DONE
//
// BoundedQueue holds no more than cap elements: Push returns false when it is full,
// and Pop returns the elements in the order they were added.
// Practices a struct with a bound and a slice.
package main_test

import "testing"

type BoundedQueue struct {
	cap   int
	items []int
}

func (q *BoundedQueue) Push(v int) bool {
	q.items = append(q.items, v)
	return true
}

func (q *BoundedQueue) Pop() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}

func TestBoundedQueue(t *testing.T) {
	q := &BoundedQueue{cap: 2}
	if !q.Push(1) || !q.Push(2) || q.Push(3) {
		t.Errorf("Push results are wrong")
	}
	if v, _ := q.Pop(); v != 1 {
		t.Errorf("Pop = %d", v)
	}
	if !q.Push(4) {
		t.Errorf("Push after Pop should succeed")
	}
}
