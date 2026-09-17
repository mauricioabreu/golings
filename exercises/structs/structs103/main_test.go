// structs103
// Make the tests pass!

// I AM NOT DONE
//
// Admin embeds User and overrides Describe, and it has to include
// the description of the user. Right now the method calls itself.
// Practices reaching a hidden method of an embedded type by its name.
package main_test

import "testing"

type User struct{ Name string }

func (u User) Describe() string { return "user " + u.Name }

type Admin struct {
	User
	Level int
}

func (a Admin) Describe() string {
	if a.Level > 100 {
		return "admin"
	}
	a.Level += 100
	return a.Describe() + " (admin)"
}

func TestAdminDescribe(t *testing.T) {
	if got := (Admin{User{"ann"}, 1}).Describe(); got != "user ann (admin)" {
		t.Errorf("Describe = %q", got)
	}
}
