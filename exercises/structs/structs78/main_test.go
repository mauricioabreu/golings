// structs78
// Make the tests pass!

// I AM NOT DONE
//
// Employee embeds Person, which has a Name field, and declares a Name field of its
// own for the job title. The function has to return the name of the person.
// A field of the outer struct hides the promoted one.
package main_test

import "testing"

type Person struct{ Name string }

type Employee struct {
	Person
	Name string
}

func personName(e Employee) string {
	return e.Name
}

func TestPersonName(t *testing.T) {
	e := Employee{Person{"Ann"}, "Engineer"}
	if personName(e) != "Ann" {
		t.Errorf("personName = %q", personName(e))
	}
}
