// anonymous functions1
// Make me compile!

package main

import "fmt"

func main() {

	func(name string) {
		fmt.Printf("Hello %s", name)
	}("Andrew")

}
