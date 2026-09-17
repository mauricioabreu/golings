// arrays77
// Make the tests pass!

// I AM NOT DONE
//
// oldest returns the name of the oldest person.
// Practices reaching the fields of the elements of an array of structs.
package main_test

import "testing"

type Person struct {
	Name string
	Age  int
}

func oldest(ps [3]Person) string {
	best := ps[0]
	for _, p := range ps {
		if p.Age > best.Age {
			best = p
		}
	}
	return ps[0].Name
}

func TestOldest(t *testing.T) {
	ps := [3]Person{{"Ann", 30}, {"Bob", 45}, {"Cid", 20}}
	if got := oldest(ps); got != "Bob" {
		t.Errorf("oldest = %s, want Bob", got)
	}
}
