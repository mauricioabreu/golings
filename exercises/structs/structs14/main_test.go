// structs14
// Make the tests pass!

// I AM NOT DONE
//
// countAdults counts the people over 17.
// Practices a range over a slice of structs.
package main_test

import "testing"

type Person struct {
	Name string
	Age  int
}

func countAdults(ps []Person) int {
	n := 0
	for _, p := range ps {
		if p.Age > 18 {
			n++
		}
	}
	return n
}

func TestCountAdults(t *testing.T) {
	if got := countAdults([]Person{{"a", 18}, {"b", 17}, {"c", 40}}); got != 2 {
		t.Errorf("countAdults = %d", got)
	}
}
