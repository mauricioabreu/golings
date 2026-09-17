// primitive_types32
// Make the tests pass!

// I AM NOT DONE
//
// encode writes a uint32 into 4 bytes in big-endian order, and decode reads it back.
// Practices the encoding/binary package.
package main_test

import (
	"encoding/binary"
	"testing"
)

func encode(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
}

func decode(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func TestEncode(t *testing.T) {
	b := encode(0x01020304)
	if b[0] != 1 || b[3] != 4 {
		t.Errorf("encode = %v, want [1 2 3 4]", b)
	}
	if got := decode([]byte{0, 0, 1, 0}); got != 256 {
		t.Errorf("decode([0 0 1 0]) = %d, want 256", got)
	}
}
