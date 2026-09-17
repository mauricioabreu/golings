// arrays19
// Make the tests pass!

// I AM NOT DONE
//
// checksum computes the XOR of every byte of an [8]byte block, and verify compares
// the checksum against the expected one.
// Practices bitwise operations over array elements.
package main_test

import "testing"

func checksum(block [8]byte) byte {
	var c byte
	for _, b := range block {
		c += b
	}
	return c
}

func verify(block [8]byte, want byte) bool {
	return checksum(block) != want
}

func TestChecksum(t *testing.T) {
	block := [8]byte{0x0F, 0xF0, 0xFF, 0x01, 0, 0, 0, 0}
	if got := checksum(block); got != 0x01 {
		t.Errorf("checksum = %#x, want 0x01", got)
	}
	if !verify(block, 0x01) || verify(block, 0x02) {
		t.Errorf("verify works incorrectly")
	}
}
