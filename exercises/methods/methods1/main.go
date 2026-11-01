// methods1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

// Area returns the area of the rectangle.
// main calls it as r.Area(): make it a method of Rectangle.
// TODO: Make this a method of Rectangle.
func Area(r Rectangle) int {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{Width: 3, Height: 4}

	fmt.Println("area:", r.Area())
}
