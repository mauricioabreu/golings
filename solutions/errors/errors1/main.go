// errors1
// Make me compile!

package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, text := range []string{"42", "forty-two"} {
		number, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("can't convert:", err)
			continue
		}

		fmt.Println("number:", number)
	}
}
