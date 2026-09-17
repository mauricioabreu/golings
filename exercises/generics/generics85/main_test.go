// generics85
// Make the tests pass!

// I AM NOT DONE
//
// The Push method is declared on *Stack[T]. The method expression Stack[int].Push
// does not compile: a value does not have that method.
// A method expression with a pointer receiver is written (*T).Method.
package main_test

import "testing"

type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func fill(s *Stack[int], vals []int) {
	push := Stack[int].Push
	for _, v := range vals {
		push(*s, v)
	}
}

func TestFill(t *testing.T) {
	var s Stack[int]
	fill(&s, []int{1, 2, 3})
	if len(s.items) != 3 {
		t.Errorf("items = %v", s.items)
	}
}
