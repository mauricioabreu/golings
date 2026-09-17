// structs99
// Make the tests pass!

// I AM NOT DONE
//
// ParseVersion parses "major.minor.patch", and Less compares versions.
// Practices a constructor from a string plus a comparison method.
package main_test

import (
	"fmt"
	"testing"
)

type Version struct{ Major, Minor, Patch int }

func ParseVersion(s string) (Version, error) {
	var v Version
	_, err := fmt.Sscanf(s, "%d.%d.%d", &v.Major, &v.Minor, &v.Patch)
	return v, err
}

func (v Version) Less(o Version) bool {
	return v.Major < o.Major || v.Minor < o.Minor || v.Patch < o.Patch
}

func TestVersion(t *testing.T) {
	a, _ := ParseVersion("1.10.0")
	b, _ := ParseVersion("2.0.1")
	c, _ := ParseVersion("1.9.9")
	if !a.Less(b) || b.Less(a) || !c.Less(a) || a.Less(c) {
		t.Errorf("Less works incorrectly: %v %v %v", a, b, c)
	}
}
