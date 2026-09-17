// switch97
// Make the tests pass!

// I AM NOT DONE
//
// route returns the name of the handler for a method and a path.
// Practices a switch on a compound string key.
package main_test

import "testing"

func route(method, path string) string {
	switch method + " " + path {
	case "GET /users":
		return "listUsers"
	case "GET /users/new":
		return "createUser"
	default:
		return "listUsers"
	}
}

func TestRoute(t *testing.T) {
	cases := []struct{ m, p, want string }{
		{"GET", "/users", "listUsers"}, {"POST", "/users", "createUser"},
		{"DELETE", "/users", "notFound"}, {"GET", "/posts", "notFound"},
	}
	for _, c := range cases {
		if got := route(c.m, c.p); got != c.want {
			t.Errorf("route(%s %s) = %s, want %s", c.m, c.p, got, c.want)
		}
	}
}
