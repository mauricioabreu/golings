// structs65
// Make the tests pass!

// I AM NOT DONE
//
// totalArea adds up the areas of different shapes through the Shape interface.
// Practices several structs satisfying one interface.
package main_test

import "testing"

type Shape interface{ Area() float64 }

type Square struct{ Side float64 }
type Triangle struct{ Base, Height float64 }

func (s Square) Area(scale float64) float64 { return s.Side * s.Side * scale }
func (t Triangle) Area() float64            { return t.Base * t.Height }

func totalArea(shapes []Shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

func TestTotalArea(t *testing.T) {
	if got := totalArea([]Shape{Square{2}, Triangle{4, 3}}); got != 10 {
		t.Errorf("totalArea = %v", got)
	}
}
