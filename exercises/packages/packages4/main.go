// packages4
// Make me compile!

// I AM NOT DONE
package main

import (
	"crypto/rand"
	"fmt"
	"math/rand/v2"
)

func main() {
	// TODO: Both imported packages are called rand. Give one of them another name with an alias.
	dice := rand.IntN(6) + 1
	fmt.Println("rolled a", dice, "of 6")

	secret := make([]byte, 4)
	rand.Read(secret)
	fmt.Println("secret has", len(secret), "bytes")
}
