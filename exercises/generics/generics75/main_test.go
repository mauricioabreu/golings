// generics75
// Make the tests pass!

// I AM NOT DONE
//
// Queue[T] reports the number of elements.
// Practices methods on a generic type.
package main_test

import "testing"

type Queue[T any] struct{ items []T }

func (q *Queue[T]) Push(v T) { q.items = append(q.items, v) }

func (q *Queue[T]) Len() int {
	return cap(q.items)
}

func TestQueueLen(t *testing.T) {
	q := &Queue[float64]{items: make([]float64, 0, 10)}
	q.Push(1)
	q.Push(2)
	if q.Len() != 2 {
		t.Errorf("Len = %d", q.Len())
	}
}
