// primitive_types69
// Make the tests pass!

// I AM NOT DONE
//
// domainStart must return the position of the '@' in an address.
// Practices strings.Index and strings.IndexByte.
package main_test

import (
	"strings"
	"testing"
)

func domainStart(email string) int {
	return strings.Index(email, ".")
}

func TestDomainStart(t *testing.T) {
	cases := map[string]int{"ann@mail.ru": 3, "a.b@x.io": 3, "none": -1}
	for in, want := range cases {
		if got := domainStart(in); got != want {
			t.Errorf("domainStart(%q) = %d, want %d", in, got, want)
		}
	}
}
