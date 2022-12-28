// anonymous functions1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func status() func() string {
	var index int
	orderStatus := map[int]string{
		1: "TO DO",
		2: "DOING",
		3: "DONE",
	}

	return func() string {
		index++
		return "What status should I return?"
	}
}

func main() {
	anonymous_func := status()
	s := anonymous_func()

	fmt.Println(s)

	if s == "DONE" {
		fmt.Println("Good Job!")
	}
}
