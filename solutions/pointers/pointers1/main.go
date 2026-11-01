// pointers1
// Make me compile!

package main

import "fmt"

func main() {
	answer := 42

	// & gives the address of a variable. The type of that address is *int.
	var pointer *int = &answer

	// * follows a pointer to the value it points to.
	var value int = *pointer

	fmt.Println("the pointer points to", value)
}
