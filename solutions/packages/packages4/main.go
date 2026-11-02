// packages4
// Make me compile!

package main

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/rand/v2"
)

func main() {
	dice := rand.IntN(6) + 1
	fmt.Println("rolled a", dice, "of 6")

	secret := make([]byte, 4)
	cryptorand.Read(secret)
	fmt.Println("secret has", len(secret), "bytes")
}
