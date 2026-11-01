// strings6
// Make me compile!

package main

import (
	"fmt"
	"strings"
)

func main() {
	names := "ada,grace,margaret,katherine"

	// strings.Split builds a slice with all the parts. strings.SplitSeq (Go 1.24)
	// returns an iterator instead, which you can range over without that slice.
	for name := range strings.SplitSeq(names, ",") {
		fmt.Println(name)
	}

	text := "first\nsecond\nthird\n"

	// strings.Lines iterates over the lines of a text, new line included.
	n := 0
	for line := range strings.Lines(text) {
		n++
		fmt.Printf("%d: %s", n, line)
	}
}
