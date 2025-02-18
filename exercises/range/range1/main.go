// range1
// Make me compile!
package main

import "fmt"

func main() {
	evenNumbers := []int{2, 4, 6, 8, 10}

	for _, v := range evenNumbers {
		fmt.Printf("%d is even\n", v)
	}
}
