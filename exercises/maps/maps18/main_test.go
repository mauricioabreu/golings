// maps18
// Make the tests pass!

// I AM NOT DONE
//
// prefixes builds a map from length to a prefix of a string, as slices of one buffer,
// and the values are appended to afterwards. Appending to one prefix spoils the others.
// Slice values in a map may share an array.
package main_test

import "testing"

func prefixes(word []byte) map[int][]byte {
	m := map[int][]byte{}
	for n := 1; n <= len(word); n++ {
		m[n] = word[:n]
	}
	return m
}

func TestPrefixes(t *testing.T) {
	word := []byte("gopher")
	m := prefixes(word)
	m[2] = append(m[2], '!')
	if string(m[3]) != "gop" || string(word) != "gopher" || string(m[2]) != "go!" {
		t.Errorf("m[2]=%q m[3]=%q word=%q", m[2], m[3], word)
	}
}
