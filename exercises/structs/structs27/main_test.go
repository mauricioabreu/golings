// structs27
// Make the tests pass!

// I AM NOT DONE
//
// Stack with the methods Push, Pop, Peek and Len. Pop and Peek on an empty stack return false.
// Practices wrapping a slice in a struct.
package main_test

import "testing"

type Stack struct{ items []string }

func (s *Stack) Push(v string) { s.items = append(s.items, v) }
func (s *Stack) Len() int      { return len(s.items) }

func (s *Stack) Peek() (string, bool) {
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

func (s *Stack) Pop() (string, bool) {
	v, ok := s.Peek()
	if ok {
		s.items = s.items[:len(s.items)-1]
	}
	return v, ok
}

func TestStack(t *testing.T) {
	var s Stack
	s.Push("a")
	s.Push("b")
	if v, _ := s.Peek(); v != "b" || s.Len() != 2 {
		t.Errorf("Peek = %q, Len = %d", v, s.Len())
	}
	if v, _ := s.Pop(); v != "b" {
		t.Errorf("Pop = %q", v)
	}
	s.Pop()
	if _, ok := s.Pop(); ok {
		t.Errorf("Pop on empty stack should fail")
	}
}
