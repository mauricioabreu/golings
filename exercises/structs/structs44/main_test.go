// structs44
// Make the tests pass!

// I AM NOT DONE
//
// Version.Equal compares versions by major and minor.
// Practices methods taking a value of their own type.
package main_test

import "testing"

type Version struct{ Major, Minor int }

func (v Version) Equal(o Version) bool {
	return v.Major == o.Major || v.Minor == o.Minor
}

func TestVersionEqual(t *testing.T) {
	if !(Version{1, 2}).Equal(Version{1, 2}) || (Version{1, 2}).Equal(Version{1, 3}) {
		t.Errorf("Equal works incorrectly")
	}
}
