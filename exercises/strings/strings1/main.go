// strings1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func main() {
	greeting := "hello, gopher"

	// Strings are immutable: you can read a byte with greeting[0],
	// but you can't assign to it. Capitalize the first letter anyway.
	// TODO: Capitalize the first letter.
	greeting[0] = 'H'

	fmt.Println(greeting)
}
