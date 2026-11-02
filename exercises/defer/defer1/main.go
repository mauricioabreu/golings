// defer1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func main() {
	fmt.Println("start")

	// TODO: The deferred function has to be called, not only declared.
	defer func() {
		fmt.Println("goodbye, printed when main returns")
	}

	fmt.Println("end")
}
