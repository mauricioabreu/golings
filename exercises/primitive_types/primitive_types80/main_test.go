// primitive_types80
// Make the tests pass!

// I AM NOT DONE
//
// letterFreq must count the latin letters ignoring case.
// Practices a counter array and turning letters into indexes.
package main_test

import "testing"

func letterFreq(s string) [26]int {
	var f [26]int
	for i := 0; i < len(s); i++ {
		c := s[i]
		f[c-'a']++
	}
	return f
}

func TestLetterFreq(t *testing.T) {
	f := letterFreq("Hello, World!")
	if f['l'-'a'] != 3 || f['h'-'a'] != 1 || f['w'-'a'] != 1 || f['o'-'a'] != 2 {
		t.Errorf("letterFreq = %v", f)
	}
}
