// variables67
// Make the tests pass!

// I AM NOT DONE
//
// toByte must clamp a value to the range 0..255.
// Right now 300 turns into 44 and -5 turns into 251.
// An integer conversion does not check the range.
package main_test

import "testing"

func toByte(v int) byte {
	return byte(v)
}

func TestToByte(t *testing.T) {
	cases := map[int]byte{-5: 0, 0: 0, 128: 128, 255: 255, 300: 255}
	for in, want := range cases {
		if got := toByte(in); got != want {
			t.Errorf("toByte(%d) = %d, want %d", in, got, want)
		}
	}
}
