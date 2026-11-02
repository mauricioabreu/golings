// packages1
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Only names that start with an upper case letter can be used from
	// another package.
	// TODO: Use the exported names of math and strings.
	fmt.Println("pi is about", math.pi)
	fmt.Println(strings.toUpper("exported names start with a capital letter"))
}
