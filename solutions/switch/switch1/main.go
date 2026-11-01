// switch1
// Make me compile!

package main

import "fmt"

func main() {
	status := "open"
	switch status {
	case "open":
		fmt.Println("status is open")
	case "closed":
		fmt.Println("status is closed")
	}
}
