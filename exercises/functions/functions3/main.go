// functions3
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func main() {
	// TODO: call_me needs a number to count up to: pass one.
	call_me()
}

func call_me(num int) {
	for n := 0; n <= num; n++ {
		fmt.Printf("Num is %d\n", n)
	}
}
