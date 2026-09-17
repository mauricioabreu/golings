// switch80
// Make the tests pass!

// I AM NOT DONE
//
// category works out the category of a value from its type. The code does not compile:
// the type switch variable is declared but used in none of the branches.
// The switch v := x.(type) form requires v to be used.
package main_test

import "testing"

func category(x interface{}) string {
	switch v := x.(type) {
	case int, float64:
		return "number"
	case string:
		return "text"
	}
	return "other"
}

func TestCategory(t *testing.T) {
	cases := []struct {
		in   interface{}
		want string
	}{{1, "number"}, {2.5, "number"}, {"a", "text"}, {true, "other"}}
	for _, c := range cases {
		if got := category(c.in); got != c.want {
			t.Errorf("category(%v) = %s, want %s", c.in, got, c.want)
		}
	}
}
