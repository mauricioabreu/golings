// maps91
// Make the tests pass!

// I AM NOT DONE
//
// sameConfig compares two configurations. The code does not compile:
// a map can only be compared with nil.
// Maps do not support ==.
package main_test

import "testing"

func sameConfig(a, b map[string]string) bool {
	return a == b
}

func TestSameConfig(t *testing.T) {
	a := map[string]string{"x": "1", "y": ""}
	if !sameConfig(a, map[string]string{"y": "", "x": "1"}) {
		t.Errorf("equal configs should match")
	}
	if sameConfig(a, map[string]string{"x": "1", "z": ""}) {
		t.Errorf("configs with different keys should not match")
	}
}
