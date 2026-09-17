// switch41
// Make the tests pass!

// I AM NOT DONE
//
// describe returns "read", "write", "read-write" or "none" for a set of flags.
// Read|Write comes out as "none": a switch on the value does not see the combination.
// Bit flags cannot be taken apart by a switch over the individual constants.
package main_test

import "testing"

const (
	Read = 1 << iota
	Write
)

func describe(flags int) string {
	switch flags {
	case Read:
		return "read"
	case Write:
		return "write"
	}
	return "none"
}

func TestDescribe(t *testing.T) {
	cases := map[int]string{0: "none", Read: "read", Write: "write", Read | Write: "read-write"}
	for in, want := range cases {
		if got := describe(in); got != want {
			t.Errorf("describe(%b) = %s, want %s", in, got, want)
		}
	}
}
