// slices1
// Make me compile!

package main

import "fmt"

func main() {
	a := make([]int, 3, 10) // play with length and capacity
	fmt.Println("length of 'a':", len(a))
	fmt.Println("capacity of 'a':", cap(a))
}
