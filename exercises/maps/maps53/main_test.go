// maps53
// Make the tests pass!

// I AM NOT DONE
//
// The indexes by email and by name hold pointers to one and the same object.
// rename has to return a new copy of the user under a different name
// without touching the object that is already indexed.
// Pointers in several maps mean one shared mutable object.
package main_test

import "testing"

type User struct{ Name, Email string }

func rename(u *User, name string) *User {
	u.Name = name
	return u
}

func TestRename(t *testing.T) {
	u := &User{"ann", "a@x"}
	byEmail := map[string]*User{u.Email: u}
	byName := map[string]*User{u.Name: u}
	r := rename(u, "anna")
	if r.Name != "anna" {
		t.Errorf("renamed = %v", r)
	}
	if byEmail["a@x"].Name != "ann" || byName["ann"].Name != "ann" {
		t.Errorf("indexed user changed: %v", byEmail["a@x"])
	}
}
