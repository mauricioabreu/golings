// switch98
// Make the tests pass!

// I AM NOT DONE
//
// kind returns "int", "string" or "other" depending on the dynamic type.
// Practices a type switch.
package main_test

import "testing"

func kind(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case []byte:
		return "string"
	default:
		return "other"
	}
}

func TestKind(t *testing.T) {
	cases := []struct {
		in   interface{}
		want string
	}{{1, "int"}, {"go", "string"}, {1.5, "other"}, {[]byte("x"), "other"}}
	for _, c := range cases {
		if got := kind(c.in); got != c.want {
			t.Errorf("kind(%v) = %s, want %s", c.in, got, c.want)
		}
	}
}
