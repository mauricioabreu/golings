// structs50
// Make the tests pass!

// I AM NOT DONE
//
// The Describe method of the embedded Animal struct is available on Dog.
// Practices method promotion through embedding.
package main_test

import "testing"

type Animal struct{ Name string }

func (a Animal) Describe() string { return "animal " + a.Name }

type Dog struct {
	A     Animal
	Breed string
}

func TestDogDescribe(t *testing.T) {
	d := Dog{Animal{"Rex"}, "husky"}
	if d.Describe() != "animal Rex" {
		t.Errorf("Describe = %q", d.Describe())
	}
}
