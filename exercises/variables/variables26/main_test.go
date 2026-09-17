// variables26
// Make the tests pass!

// I AM NOT DONE
//
// revoke must clear a flag whether or not it was set.
// Calling it twice currently turns the flag back on.
// Practices the bit clear operator &^ and how it differs from XOR.
package main_test

import "testing"

const (
	FlagRead uint8 = 1 << iota
	FlagWrite
	FlagAdmin
)

func revoke(perms, flag uint8) uint8 {
	return perms ^ flag
}

func TestRevoke(t *testing.T) {
	p := FlagRead | FlagAdmin
	p = revoke(p, FlagAdmin)
	if p != FlagRead {
		t.Errorf("after first revoke perms = %03b, want %03b", p, FlagRead)
	}
	p = revoke(p, FlagAdmin)
	if p != FlagRead {
		t.Errorf("after second revoke perms = %03b, want %03b", p, FlagRead)
	}
}
