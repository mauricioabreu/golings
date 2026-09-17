// structs93
// Make the tests pass!

// I AM NOT DONE
//
// Rename changes the name of a user.
// Practices a method with a pointer receiver.
package main_test

import "testing"

type User struct{ Name string }

func (u *User) Rename(name string) {
	u.Name = u.Name
}

func TestRename(t *testing.T) {
	u := &User{"ann"}
	u.Rename("anna")
	if u.Name != "anna" {
		t.Errorf("Name = %q", u.Name)
	}
}
