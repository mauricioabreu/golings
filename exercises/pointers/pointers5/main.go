// pointers5
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	// new(T) allocates a zero value of type T and returns a pointer to it.
	count := new(int)
	// TODO: Add one to the number the pointer points to.
	count++
	fmt.Println("count:", *count)

	origin := new(Point)
	origin.X = 3
	fmt.Println("origin:", *origin)
}
