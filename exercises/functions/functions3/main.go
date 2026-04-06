// functions3
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func main() {
	callMe()
}

func callMe(num int) {
	for n := 0; n <= num; n++ {
		fmt.Printf("Num is %d\n", n)
	}
}
