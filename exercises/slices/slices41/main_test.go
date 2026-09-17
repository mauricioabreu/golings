// slices41
// Make the tests pass!

// I AM NOT DONE
//
// addAndRename adds a user and then changes the name of the first one through
// a pointer taken earlier. The change is lost.
// append may move the elements to a new array, leaving old pointers stale.
package main_test

import "testing"

type User struct{ Name string }

func addAndRename(users []User) []User {
	first := &users[0]
	users = append(users, User{"new"})
	first.Name = "admin"
	return users
}

func TestAddAndRename(t *testing.T) {
	users := []User{{"ann"}}
	users = addAndRename(users)
	if users[0].Name != "admin" || users[1].Name != "new" {
		t.Errorf("users = %v", users)
	}
}
