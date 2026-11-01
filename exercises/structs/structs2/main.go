// structs2
// Make me compile!
//
// I AM NOT DONE
package main

import "fmt"

type Person struct {
	// TODO: Don't just create the phone field here: embed a new struct, ContactDetails.
	name string
	age int
}

func main() {
	// contactDetails := ContactDetails{}
	person := Person{name: "John", age: 32}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
