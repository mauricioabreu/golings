// switch1
// Make me compile!

package main

import "fmt"

func main() {
	status := "open"
	switch {
	case status == "open":
		fmt.Println("status is open")
	case status == "closed":
		fmt.Println("status is closed")
	}
}
