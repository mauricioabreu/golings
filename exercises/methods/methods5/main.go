// methods5
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

// Incrementer is any type that has an Increment method. A variable of this
// type can hold any value that has the method in its method set.
type Incrementer interface {
	Increment()
}

func main() {
	c := Counter{}

	// TODO: Put into i a value whose type has the Increment method.
	var i Incrementer = c
	i.Increment()

	fmt.Println("count:", c.count)
}
