// algorithms71
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a stack with extra state. Implement a stack with Push, Pop, Top,
// and Min operations, each in O(1). Operations on an empty stack are never
// called.
// Expected asymptotics: O(1) per operation, O(n) space.
package main_test

import "testing"

type MinStack struct {
	vals, mins []int
}

func (s *MinStack) Push(v int) {
}

func (s *MinStack) Pop() {
}

func (s *MinStack) Top() int {
	return 0
}

func (s *MinStack) Min() int {
	return 0
}

func TestMinStack(t *testing.T) {
	var s MinStack
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if s.Min() != -3 {
		t.Errorf("Min = %d, want -3", s.Min())
	}
	s.Pop()
	if s.Top() != 0 || s.Min() != -2 {
		t.Errorf("Top = %d, Min = %d", s.Top(), s.Min())
	}
	s.Push(5)
	s.Push(-2)
	s.Pop()
	if s.Min() != -2 || s.Top() != 5 {
		t.Errorf("after pushes Top = %d, Min = %d", s.Top(), s.Min())
	}
}
