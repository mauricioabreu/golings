// generics1
// Make me compile!
package main

import "fmt"

func main() {
	print("Hello, World!")
	print(42)
}

func print(value any) {
	fmt.Println(value)
}
