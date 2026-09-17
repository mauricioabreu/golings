// maps35
// Make the tests pass!

// I AM NOT DONE
//
// emailOf returns the email of a user from a map of structs.
// Practices reaching a field of a map value.
package main_test

import "testing"

type User struct {
	Name  string
	Email string
}

func emailOf(users map[int]User, id int) string {
	return users[id].Name
}

func TestEmailOf(t *testing.T) {
	users := map[int]User{7: {"ann", "ann@x.io"}}
	if emailOf(users, 7) != "ann@x.io" || emailOf(users, 8) != "" {
		t.Errorf("emailOf works incorrectly")
	}
}
