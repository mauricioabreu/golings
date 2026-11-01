// methods4
// Make me compile!

package main

import "fmt"

// Celsius is a temperature. You can define methods on types you declare,
// not only on structs.
type Celsius float64

// ToFahrenheit converts the temperature to degrees Fahrenheit.
func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func main() {
	boiling := Celsius(100)

	fmt.Println("water boils at", boiling.ToFahrenheit(), "degrees Fahrenheit")
}
