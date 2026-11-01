// pointers1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func main() {
	answer := 42

	// & gives the address of a variable. The type of that address is *int.
	// TODO: Make the pointer point at answer.
	var pointer *int = answer

	// * follows a pointer to the value it points to.
	// TODO: Read the value the pointer points to.
	var value int = pointer

	fmt.Println("the pointer points to", value)
}
