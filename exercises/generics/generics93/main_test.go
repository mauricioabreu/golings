// generics93
// Make the tests pass!

// I AM NOT DONE
//
// The Queue.Push method is declared with a Queue receiver, without the [T].
// The code does not compile.
// The receiver of a method on a generic type has to list the parameters.
package main_test

import "testing"

type Queue[T any] struct{ items []T }

func (q *Queue) Push(v T) { q.items = append(q.items, v) }

func (q *Queue[T]) Len() int { return len(q.items) }

func TestQueuePush(t *testing.T) {
	var q Queue[string]
	q.Push("a")
	if q.Len() != 1 {
		t.Errorf("Len = %d", q.Len())
	}
}
