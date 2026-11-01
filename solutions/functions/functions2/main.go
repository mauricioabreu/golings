// functions2
// Make me compile!

package main

import "fmt"

func main() {
	callMe(10)
}

func callMe(num int) {
	for n := 0; n <= num; n++ {
		fmt.Printf("Num is %d\n", n)
	}
}
