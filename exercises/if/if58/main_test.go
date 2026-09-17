// if58
// Make the tests pass!

// I AM NOT DONE
//
// describe must return the string in upper case when the value is a string,
// and "not a string" otherwise. Right now a number makes it panic.
// The v, ok := x.(T) form does not panic on a type mismatch.
package main_test

import (
	"strings"
	"testing"
)

func describe(v interface{}) string {
	if s := v.(string); s != "" {
		return strings.ToUpper(s)
	}
	return "not a string"
}

func TestDescribe(t *testing.T) {
	cases := []struct {
		in   interface{}
		want string
	}{{"go", "GO"}, {42, "not a string"}, {nil, "not a string"}, {"", ""}}
	for _, c := range cases {
		if got := describe(c.in); got != c.want {
			t.Errorf("describe(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
