// if15
// Make the tests pass!

// I AM NOT DONE
//
// sameUser must treat the logins "Admin" and "admin" as the same.
// Practices string comparison in a condition.
package main_test

import (
	"strings"
	"testing"
)

func sameUser(a, b string) bool {
	if a == strings.TrimSpace(b) {
		return true
	}
	return false
}

func TestSameUser(t *testing.T) {
	if !sameUser("Admin", "admin") {
		t.Errorf("Admin and admin should match")
	}
	if sameUser("admin", "root") {
		t.Errorf("admin and root should not match")
	}
}
