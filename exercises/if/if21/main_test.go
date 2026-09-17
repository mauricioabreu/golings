// if21
// Make the tests pass!

// I AM NOT DONE
//
// login returns "no such user", "locked", "wrong password" or "welcome".
// A locked user must not be told whether the password was right.
// Practices the order of the checks and comma-ok on a map.
package main_test

import "testing"

type account struct {
	password string
	locked   bool
}

var users = map[string]account{
	"ann": {password: "secret"},
	"bob": {password: "qwerty", locked: true},
}

func login(name, pw string) string {
	u := users[name]
	if u.password != pw {
		return "wrong password"
	}
	if u.locked {
		return "locked"
	}
	return "welcome"
}

func TestLogin(t *testing.T) {
	cases := []struct{ name, pw, want string }{
		{"ann", "secret", "welcome"}, {"ann", "x", "wrong password"},
		{"bob", "x", "locked"}, {"bob", "qwerty", "locked"}, {"eve", "", "no such user"},
	}
	for _, c := range cases {
		if got := login(c.name, c.pw); got != c.want {
			t.Errorf("login(%s, %s) = %q, want %q", c.name, c.pw, got, c.want)
		}
	}
}
