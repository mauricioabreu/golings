// functions2
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func main() {
	callMe(10)
}

// TODO: Give the parameter a type.
func callMe(num) {
	for n := 0; n <= num; n++ {
		fmt.Printf("Num is %d\n", n)
	}
}
