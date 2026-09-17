// generics65
// Make the tests pass!

// I AM NOT DONE
//
// Stack[T] offers Peek and Pop, both returning (T, bool).
// Practices the zero value of T for an empty stack.
package main_test

import "testing"

type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Peek() (T, bool) {
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Pop() (T, bool) {
	v, ok := s.Peek()
	if ok {
		s.items = s.items[:len(s.items)-1]
	}
	return v, ok
}

func TestStack(t *testing.T) {
	var s Stack[float64]
	if _, ok := s.Pop(); ok {
		t.Errorf("Pop on empty stack should fail")
	}
	s.Push(1.5)
	if v, ok := s.Peek(); !ok || v != 1.5 || len(s.items) != 1 {
		t.Errorf("Peek = %v, %v", v, ok)
	}
}
