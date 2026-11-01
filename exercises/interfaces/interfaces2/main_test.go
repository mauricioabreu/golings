// interfaces2
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

// totalArea adds up the areas of all the shapes.
func totalArea(shapes []Shape) float64 {
	return 0
}

// largest returns the shape with the biggest area, or nil if there are no shapes.
func largest(shapes []Shape) Shape {
	return nil
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{Rectangle{Width: 2, Height: 3}, Triangle{Base: 4, Height: 5}}

	if got := totalArea(shapes); got != 16 {
		t.Errorf("totalArea = %v, want 16", got)
	}
}

func TestLargest(t *testing.T) {
	small := Rectangle{Width: 1, Height: 1}
	big := Triangle{Base: 10, Height: 10}

	if got := largest([]Shape{small, big, Rectangle{Width: 2, Height: 3}}); got != big {
		t.Errorf("largest = %v, want %v", got, big)
	}
	if got := largest(nil); got != nil {
		t.Errorf("largest of no shapes = %v, want nil", got)
	}
}
