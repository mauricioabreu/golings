// strings1
// Make me compile!

package main

import "fmt"

func main() {
	greeting := "hello, gopher"

	// Strings are immutable: you can read a byte with greeting[0],
	// but you can't assign to it. Capitalize the first letter anyway.
	letters := []byte(greeting)
	letters[0] = 'H'
	greeting = string(letters)

	fmt.Println(greeting)
}
