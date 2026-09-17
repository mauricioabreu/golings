// variables25
// Make the tests pass!

// I AM NOT DONE
//
// Permissions are stored as a set of bit flags.
// can must report whether a given flag is set.
// Practices 1 << iota and the bitwise operators & and |.
package main_test

import "testing"

type Perm uint8

const (
	Read Perm = iota
	Write
	Exec
)

func can(p, flag Perm) bool {
	return p|flag != 0
}

func TestCan(t *testing.T) {
	p := Read | Exec
	if !can(p, Read) || !can(p, Exec) {
		t.Errorf("Read|Exec should allow Read and Exec")
	}
	if can(p, Write) {
		t.Errorf("Read|Exec should not allow Write")
	}
}
