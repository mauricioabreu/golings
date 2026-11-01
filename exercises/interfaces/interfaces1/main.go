// interfaces1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// Shape is satisfied by any type that has an Area method.
// There is no "implements": having the method is enough.
type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Triangle struct {
	Base, Height float64
}

func printArea(s Shape) {
	fmt.Printf("area: %.1f\n", s.Area())
}

func main() {
	printArea(Square{Side: 2})
	printArea(Triangle{Base: 3, Height: 4})
}
