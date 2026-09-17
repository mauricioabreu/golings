// anonymous_functions19
// Make the tests pass!

// I AM NOT DONE
//
// idGen(prefix) returns a literal handing out "prefix-0001", "prefix-0002" and so on.
// Practices a closure plus zero padded formatting.
package main_test

import (
	"fmt"
	"testing"
)

func idGen(prefix string) func() string {
	n := 0
	return func() string {
		return fmt.Sprintf("%s-%d", prefix, n)
	}
}

func TestIDGen(t *testing.T) {
	next := idGen("ord")
	next()
	if got := next(); got != "ord-0002" {
		t.Errorf("second id = %q", got)
	}
	if got := idGen("inv")(); got != "inv-0001" {
		t.Errorf("independent generator = %q", got)
	}
}
