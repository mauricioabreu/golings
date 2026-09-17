// structs28
// Make the tests pass!

// I AM NOT DONE
//
// area takes a shape through an interface, and the shapes are passed as pointers.
// A type assertion to a Square value never succeeds.
// The interface holds a *Square, not a Square.
package main_test

import "testing"

type Shape interface{ Name() string }

type Square struct{ Side float64 }

func (s *Square) Name() string { return "square" }

func area(s Shape) float64 {
	if sq, ok := s.(Square); ok {
		return sq.Side * sq.Side
	}
	return 0
}

func TestArea(t *testing.T) {
	if got := area(&Square{3}); got != 9 {
		t.Errorf("area = %v, want 9", got)
	}
}
