// anonymous_functions82
// Make the tests pass!

// I AM NOT DONE
//
// The greet function is decorated by reassigning the variable with a literal that
// calls greet. The literal ends up calling itself instead of the original function.
// A literal captures the variable, not the value it held.
package main_test

import (
	"strings"
	"testing"
)

func decorated() func(string) string {
	greet := func(name string) string { return "hello, " + name }
	depth := 0
	greet = func(name string) string {
		depth++
		if depth > 3 {
			return "recursion"
		}
		return strings.ToUpper(greet(name))
	}
	return greet
}

func TestDecorated(t *testing.T) {
	if got := decorated()("go"); got != "HELLO, GO" {
		t.Errorf("decorated = %q", got)
	}
}
