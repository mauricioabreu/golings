// defer1
// Make me compile!

package main

import "fmt"

func main() {
	fmt.Println("start")

	defer func() {
		fmt.Println("goodbye, printed when main returns")
	}()

	fmt.Println("end")
}
