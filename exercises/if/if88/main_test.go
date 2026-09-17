// if88
// Make the tests pass!

// I AM NOT DONE
//
// hasAll must return true only when EVERY bit of mask is set.
// Right now a single matching bit is enough.
// Practices the difference between flags&mask != 0 and flags&mask == mask.
package main_test

import "testing"

const (
	Read uint8 = 1 << iota
	Write
	Exec
)

func hasAll(flags, mask uint8) bool {
	if flags&mask != 0 {
		return true
	}
	return false
}

func TestHasAll(t *testing.T) {
	if !hasAll(Read|Write|Exec, Read|Write) {
		t.Errorf("rwx has rw")
	}
	if hasAll(Read, Read|Write) {
		t.Errorf("r does not have rw")
	}
	if hasAll(Exec, Read|Write) {
		t.Errorf("x does not have rw")
	}
}
