// variables76
// Make the tests pass!

// I AM NOT DONE
//
// checksum must return the sum of all the bytes of a string.
// For long strings the sum comes out suspiciously small.
// Practices byte (uint8) overflow while accumulating.
package main_test

import "testing"

func checksum(s string) int {
	var sum byte
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return int(sum)
}

func TestChecksum(t *testing.T) {
	if got := checksum("abc"); got != 294 {
		t.Errorf("checksum(abc) = %d, want 294", got)
	}
	if got := checksum("zzzz"); got != 488 {
		t.Errorf("checksum(zzzz) = %d, want 488", got)
	}
}
