// maps81
// Make the tests pass!

// I AM NOT DONE
//
// dedupeByEmail drops the users with a repeated email, keeping the first one.
// Practices a map used as a set over a struct field.
package main_test

import (
	"reflect"
	"testing"
)

type User struct{ Name, Email string }

func dedupeByEmail(users []User) []User {
	seen := map[string]bool{}
	var out []User
	for _, u := range users {
		if !seen[u.Name] {
			out = append(out, u)
		}
	}
	return out
}

func TestDedupeByEmail(t *testing.T) {
	got := dedupeByEmail([]User{{"ann", "a@x"}, {"anna", "a@x"}, {"bob", "b@x"}})
	if !reflect.DeepEqual(got, []User{{"ann", "a@x"}, {"bob", "b@x"}}) {
		t.Errorf("dedupeByEmail = %v", got)
	}
}
