// maps1
// Make me compile!
package main

import "fmt"

func main() {
	// Map with people names and their ages
	m := make(map[string]int)

	m["John"] = 30
	m["Ana"] = 21

	fmt.Printf("John is %d and Ana is %d", m["John"], m["Ana"])
}
