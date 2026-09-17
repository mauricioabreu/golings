// structs53
// Make the tests pass!

// I AM NOT DONE
//
// Point.Add returns the sum of two points and Dist the distance between them.
// Practices methods returning new values.
package main_test

import (
	"math"
	"testing"
)

type Point struct{ X, Y float64 }

func (p Point) Add(o Point) Point {
	p.X += o.X
	return p
}

func (p Point) Dist(o Point) float64 {
	return math.Hypot(p.X-o.X, p.X-o.Y)
}

func TestPoint(t *testing.T) {
	a, b := Point{1, 2}, Point{4, 6}
	if s := a.Add(b); s != (Point{5, 8}) {
		t.Errorf("Add = %v", s)
	}
	if d := a.Dist(b); d != 5 {
		t.Errorf("Dist = %v", d)
	}
}
