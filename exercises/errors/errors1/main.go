// errors1
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, text := range []string{"42", "forty-two"} {
		// TODO: strconv.Atoi returns two values, the number and an error. Receive both, and when there is an error print it and move on to the next text.
		number := strconv.Atoi(text)

		fmt.Println("number:", number)
	}
}
