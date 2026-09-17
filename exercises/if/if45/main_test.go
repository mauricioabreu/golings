// if45
// Make the tests pass!

// I AM NOT DONE
//
// nameOf must return the name of the user, or "anonymous" for nil.
// Practices checking a pointer before reaching for a field.
package main_test

import "testing"

type User struct{ Name string }

func nameOf(u *User) string {
	if u != nil {
		return "anonymous"
	}
	return u.Name
}

func TestNameOf(t *testing.T) {
	if got := nameOf(nil); got != "anonymous" {
		t.Errorf("nameOf(nil) = %q", got)
	}
	if got := nameOf(&User{"ann"}); got != "ann" {
		t.Errorf("nameOf(ann) = %q", got)
	}
}
