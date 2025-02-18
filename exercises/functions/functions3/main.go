// functions3
// Make me compile!

package main

import "fmt"

func main() {
	call_me(0)
}

func call_me(num int) {
	for n := 0; n <= num; n++ {
		fmt.Printf("Num is %d\n", n)
	}
}
