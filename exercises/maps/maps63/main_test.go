// maps63
// Make the tests pass!

// I AM NOT DONE
//
// countLines counts equal lines held as [][]byte. The code does not compile:
// a []byte cannot be used as a map key.
// A map key has to be of a comparable type.
package main_test

import (
	"bytes"
	"testing"
)

func countLines(data []byte) map[string]int {
	m := map[string]int{}
	for _, line := range bytes.Split(data, []byte("\n")) {
		m[line]++
	}
	return m
}

func TestCountLines(t *testing.T) {
	m := countLines([]byte("a\nb\na"))
	if m["a"] != 2 || m["b"] != 1 {
		t.Errorf("countLines = %v", m)
	}
}
