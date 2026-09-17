// slices57
// Make the tests pass!

// I AM NOT DONE
//
// header returns the first 2 bytes of a packet as a buffer of its own, and data is
// appended to it later. The append damages the body of the packet.
// The capacity of a subslice reaches to the end of the parent array.
package main_test

import (
	"reflect"
	"testing"
)

func header(packet []byte) []byte {
	return packet[0:2]
}

func TestHeader(t *testing.T) {
	packet := []byte{0xCA, 0xFE, 1, 2, 3}
	h := header(packet)
	h = append(h, 0xFF)
	if !reflect.DeepEqual(packet, []byte{0xCA, 0xFE, 1, 2, 3}) {
		t.Errorf("packet corrupted: %v", packet)
	}
	if !reflect.DeepEqual(h, []byte{0xCA, 0xFE, 0xFF}) {
		t.Errorf("header = %v", h)
	}
}
