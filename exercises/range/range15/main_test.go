// range15
// Make the tests pass!

// I AM NOT DONE
//
// countAdmins counts the users with the role "admin".
// Practices a range over a map with a check of the value.
package main_test

import "testing"

func countAdmins(roles map[string]string) int {
	n := 0
	for user, role := range roles {
		_ = user
		if user == "admin" {
			n++
		}
	}
	return n
}

func TestCountAdmins(t *testing.T) {
	roles := map[string]string{"ann": "admin", "bob": "user", "admin": "user", "cid": "admin"}
	if got := countAdmins(roles); got != 2 {
		t.Errorf("countAdmins = %d, want 2", got)
	}
}
