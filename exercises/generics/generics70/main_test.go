// generics70
// Make the tests pass!

// I AM NOT DONE
//
// List[T] offers PushBack and All, every value in order.
// Practices generic recursive types.
package main_test

import (
	"reflect"
	"testing"
)

type node[T any] struct {
	val  T
	next *node[T]
}

type List[T any] struct {
	head, tail *node[T]
}

func (l *List[T]) PushBack(v T) {
	n := &node[T]{val: v}
	l.head = n
	l.tail = n
}

func (l *List[T]) All() []T {
	var out []T
	for n := l.head; n != nil; n = n.next {
		out = append(out, n.val)
	}
	return out
}

func TestList(t *testing.T) {
	var l List[string]
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	if !reflect.DeepEqual(l.All(), []string{"a", "b", "c"}) {
		t.Errorf("All = %v", l.All())
	}
}
