// maps37
// Make the tests pass!

// I AM NOT DONE
//
// The registry keeps roles per user. A lookup with the "same" user,
// a new struct with the same ID, finds nothing.
// A pointer key compares by address rather than by contents.
package main_test

import "testing"

type User struct {
	ID   int
	Name string
}

type Roles map[*User]string

func (r Roles) Set(u User, role string) { r[&u] = role }
func (r Roles) Get(u User) string       { return r[&u] }

func TestRoles(t *testing.T) {
	r := Roles{}
	r.Set(User{ID: 1, Name: "ann"}, "admin")
	if got := r.Get(User{ID: 1, Name: "ann"}); got != "admin" {
		t.Errorf("Get = %q, want admin", got)
	}
}
