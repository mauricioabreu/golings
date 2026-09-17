// arrays95
// Make the tests pass!

// I AM NOT DONE
//
// Stack holds up to 4 elements. Push on a full stack and Pop on an empty one
// return an error.
// Practices an array with a separate fill counter.
package main_test

import (
	"errors"
	"testing"
)

type Stack struct {
	data [4]int
	n    int
}

func (s *Stack) Push(v int) error {
	s.data[s.n] = v
	s.n++
	return nil
}

func (s *Stack) Pop() (int, error) {
	s.n--
	return s.data[s.n], nil
}

func TestStack(t *testing.T) {
	_ = errors.New
	var s Stack
	for i := 1; i <= 4; i++ {
		if err := s.Push(i); err != nil {
			t.Fatalf("Push(%d) = %v", i, err)
		}
	}
	if err := s.Push(5); err == nil {
		t.Errorf("Push to full stack should fail")
	}
	for want := 4; want >= 1; want-- {
		if v, err := s.Pop(); err != nil || v != want {
			t.Errorf("Pop = %d, %v; want %d", v, err, want)
		}
	}
	if _, err := s.Pop(); err == nil {
		t.Errorf("Pop from empty stack should fail")
	}
}
