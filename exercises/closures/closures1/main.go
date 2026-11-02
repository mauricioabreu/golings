// closures1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// newCounter returns a function that counts: 1 the first time it is called,
// 2 the second time, and so on. The function remembers count between calls.
func newCounter() func() int {
	count := 0

	// TODO: The function has to return the new count.
	return func() {
		count++
	}
}

func main() {
	next := newCounter()

	fmt.Println(next(), next(), next())
}
