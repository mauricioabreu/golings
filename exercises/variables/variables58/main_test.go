// variables58
// Make the tests pass!

// I AM NOT DONE
//
// findUser must return a pointer to the user IN THE SLICE,
// so that a change made through the pointer shows up in the original data.
// The range variable is a copy of the element, not the element itself.
package main_test

import "testing"

type User struct {
	ID   int
	Name string
}

func findUser(users []User, id int) *User {
	for _, u := range users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}

func TestFindUser(t *testing.T) {
	users := []User{{1, "ann"}, {2, "bob"}}
	u := findUser(users, 2)
	if u == nil {
		t.Fatalf("findUser(2) = nil")
	}
	u.Name = "robert"
	if users[1].Name != "robert" {
		t.Errorf("users[1].Name = %q, want robert", users[1].Name)
	}
	if findUser(users, 3) != nil {
		t.Errorf("findUser(3) should be nil")
	}
}
