// packages1
// Make me compile!

package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Only names that start with an upper case letter can be used from
	// another package.
	fmt.Println("pi is about", math.Pi)
	fmt.Println(strings.ToUpper("exported names start with a capital letter"))
}
