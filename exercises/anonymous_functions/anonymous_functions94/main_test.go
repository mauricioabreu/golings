// anonymous_functions94
// Make the tests pass!

// I AM NOT DONE
//
// Meter keeps a counter literal in a field. A copy of a Meter goes on counting
// together with the original, because the literal captured shared state.
// Copying a struct copies the function value, not its state.
package main_test

import "testing"

type Meter struct {
	Next func() int
	cur  func() int
}

func NewMeter(start int) Meter {
	n := start
	return Meter{
		Next: func() int { n++; return n },
		cur:  func() int { return n },
	}
}

func (m Meter) Clone() Meter {
	return m
}

func TestMeterClone(t *testing.T) {
	a := NewMeter(0)
	a.Next()
	b := a.Clone()
	b.Next()
	b.Next()
	if got := a.Next(); got != 2 {
		t.Errorf("original counter = %d, want 2", got)
	}
	if got := b.Next(); got != 4 {
		t.Errorf("clone counter = %d, want 4", got)
	}
}
