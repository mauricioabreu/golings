// switch62
// Make the tests pass!

// I AM NOT DONE
//
// area computes the area of a shape. The shapes are passed as pointers,
// while the type switch lists value types.
// *Circle and Circle are different types to a type switch.
package main_test

import "testing"

type Circle struct{ R float64 }
type Square struct{ Side float64 }

func area(shape interface{}) float64 {
	switch s := shape.(type) {
	case Circle:
		return 3 * s.R * s.R
	case Square:
		return s.Side * s.Side
	}
	return -1
}

func TestArea(t *testing.T) {
	if got := area(&Circle{R: 2}); got != 12 {
		t.Errorf("area(&Circle{2}) = %v, want 12", got)
	}
	if got := area(&Square{Side: 3}); got != 9 {
		t.Errorf("area(&Square{3}) = %v, want 9", got)
	}
}
