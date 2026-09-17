// algorithms103
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: a stack of stacks by frequency. Push adds a value; Pop removes
// and returns the most frequent value; on ties — the one added last.
// Expected asymptotics: O(1) per operation, O(n) space.
package main_test

import "testing"

type FreqStack struct {
	freq    map[int]int
	groups  map[int][]int
	maxFreq int
}

func NewFreqStack() *FreqStack {
	return &FreqStack{freq: map[int]int{}, groups: map[int][]int{}}
}

func (s *FreqStack) Push(v int) {
}

func (s *FreqStack) Pop() int {
	return 0
}

func TestFreqStack(t *testing.T) {
	s := NewFreqStack()
	for _, v := range []int{5, 7, 5, 7, 4, 5} {
		s.Push(v)
	}
	want := []int{5, 7, 5, 4, 7, 5}
	for i, w := range want {
		if got := s.Pop(); got != w {
			t.Errorf("pop %d = %d, want %d", i, got, w)
		}
	}
}
