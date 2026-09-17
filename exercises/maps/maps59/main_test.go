// maps59
// Make the tests pass!

// I AM NOT DONE
//
// Visits are counted under the key {city, *User}. Users that are equal
// in content count as different ones.
// A struct key compares its pointer fields by address.
package main_test

import "testing"

type Person struct{ ID int }

type visitKey struct {
	City   string
	Person *Person
}

func key(city string, p Person) visitKey { return visitKey{city, &p} }

func TestVisits(t *testing.T) {
	visits := map[visitKey]int{}
	visits[key("Kazan", Person{ID: 7})]++
	visits[key("Kazan", Person{ID: 7})]++
	if len(visits) != 1 || visits[key("Kazan", Person{ID: 7})] != 2 {
		t.Errorf("visits = %v", visits)
	}
}
