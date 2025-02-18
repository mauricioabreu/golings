// structs3
// Make me compile!
package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func main() {
	person := Person{firstName: "Maurício", lastName: "Antunes"}
	fmt.Printf("Person full name is: %s\n", person.FullName()) // here it must output Person full name is: Maurício Antunes
}
