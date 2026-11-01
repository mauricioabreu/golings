// methods4
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// Celsius is a temperature. You can define methods on types you declare,
// not only on structs.
type Celsius float64

// ToFahrenheit converts the temperature to degrees Fahrenheit.
// TODO: Declare the method on a type you defined.
func (c float64) ToFahrenheit() float64 {
	return c*9/5 + 32
}

func main() {
	boiling := Celsius(100)

	fmt.Println("water boils at", boiling.ToFahrenheit(), "degrees Fahrenheit")
}
