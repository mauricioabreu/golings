// slices86
// Make the tests pass!

// I AM NOT DONE
//
// Stack implements Push and Pop on top of a slice. Pop on an empty stack returns false.
// Practices append and the slice [:len-1].
package main_test

import "testing"

type Stack struct{ items []int }

func (s *Stack) Push(v int) { s.items = append(s.items, v) }

func (s *Stack) Pop() (int, bool) {
	v := s.items[0]
	s.items = s.items[1:]
	return v, true
}

func TestStack(t *testing.T) {
	var s Stack
	s.Push(1)
	s.Push(2)
	if v, ok := s.Pop(); !ok || v != 2 {
		t.Errorf("Pop = %d, %v; want 2", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != 1 {
		t.Errorf("Pop = %d, %v; want 1", v, ok)
	}
	if _, ok := s.Pop(); ok {
		t.Errorf("Pop on empty stack should fail")
	}
}
