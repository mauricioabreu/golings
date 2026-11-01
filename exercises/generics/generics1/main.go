// generics1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func main() {
	print("Hello, World!")
	print(42)
}

// TODO: Make this function generic, so that it accepts a value of any type.
func print(value) {
	fmt.Println(value)
}
