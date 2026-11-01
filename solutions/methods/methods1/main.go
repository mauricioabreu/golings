// methods1
// Make me compile!

package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

// Area returns the area of the rectangle.
// main calls it as r.Area(): make it a method of Rectangle.
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{Width: 3, Height: 4}

	fmt.Println("area:", r.Area())
}
