// structs64
// Make the tests pass!

// I AM NOT DONE
//
// setCity changes the city in the address of a user through a pointer to the user.
// Practices changing a nested field.
package main_test

import "testing"

type Address struct{ City string }

type User struct {
	Name string
	Addr Address
}

func setCity(u *User, city string) {
	u.Name = city
}

func TestSetCity(t *testing.T) {
	u := User{Name: "ann"}
	setCity(&u, "Perm")
	if u.Addr.City != "Perm" || u.Name != "ann" {
		t.Errorf("u = %+v", u)
	}
}
