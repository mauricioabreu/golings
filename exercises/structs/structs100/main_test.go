// structs100
// Make the tests pass!

// I AM NOT DONE
//
// An Email field was added to the User struct between Name and Age.
// The positional literals stopped compiling, or took the wrong values.
// Literals with field names survive changes to a struct.
package main_test

import "testing"

type User struct {
	Name  string
	Email string
	Age   int
}

func defaultUser() User {
	return User{"guest", 18}
}

func TestDefaultUser(t *testing.T) {
	u := defaultUser()
	if u.Name != "guest" || u.Age != 18 || u.Email != "" {
		t.Errorf("defaultUser = %+v", u)
	}
}
