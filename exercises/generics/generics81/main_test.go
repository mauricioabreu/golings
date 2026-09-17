// generics81
// Make the tests pass!

// I AM NOT DONE
//
// PQ[T] is a priority queue on container/heap with a comparison function.
// Practices a generic type implementing heap.Interface.
package main_test

import (
	"container/heap"
	"reflect"
	"testing"
)

type PQ[T any] struct {
	items []T
	less  func(a, b T) bool
}

func (q *PQ[T]) Len() int           { return len(q.items) }
func (q *PQ[T]) Less(i, j int) bool { return q.less(q.items[j], q.items[i]) }
func (q *PQ[T]) Swap(i, j int)      { q.items[i], q.items[j] = q.items[j], q.items[i] }
func (q *PQ[T]) Push(x any)         { q.items = append(q.items, x.(T)) }

func (q *PQ[T]) Pop() any {
	v := q.items[0]
	q.items = q.items[1:]
	return v
}

func TestPQ(t *testing.T) {
	q := &PQ[string]{less: func(a, b string) bool { return len(a) < len(b) }}
	for _, s := range []string{"ccc", "a", "bbbb", "bb"} {
		heap.Push(q, s)
	}
	var got []string
	for q.Len() > 0 {
		got = append(got, heap.Pop(q).(string))
	}
	if !reflect.DeepEqual(got, []string{"a", "bb", "ccc", "bbbb"}) {
		t.Errorf("order = %v", got)
	}
}
