// switch88
// Make the tests pass!

// I AM NOT DONE
//
// httpMethod returns the action for a method, or "unsupported" for the rest.
// Practices the default branch.
package main_test

import "testing"

func httpMethod(m string) string {
	switch m {
	case "GET":
		return "read"
	case "POST":
		return "create"
	default:
		return "read"
	}
}

func TestHTTPMethod(t *testing.T) {
	cases := map[string]string{"GET": "read", "POST": "create", "PATCH": "unsupported"}
	for in, want := range cases {
		if got := httpMethod(in); got != want {
			t.Errorf("httpMethod(%s) = %s, want %s", in, got, want)
		}
	}
}
