// switch35
// Make the tests pass!

// I AM NOT DONE
//
// describe must return "nil" for a nil interface as well as for a nil *User pointer.
// A nil pointer in an interface lands in the branch of its own type, not in case nil.
package main_test

import "testing"

type User struct{ Name string }

func describe(v interface{}) string {
	switch u := v.(type) {
	case nil:
		return "nil"
	case *User:
		return "user " + u.Name
	}
	return "other"
}

func TestDescribe(t *testing.T) {
	var none *User
	cases := []struct {
		in   interface{}
		want string
	}{{nil, "nil"}, {none, "nil"}, {&User{"ann"}, "user ann"}, {5, "other"}}
	for _, c := range cases {
		if got := describe(c.in); got != c.want {
			t.Errorf("describe(%#v) = %s, want %s", c.in, got, c.want)
		}
	}
}
