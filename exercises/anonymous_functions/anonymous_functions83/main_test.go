// anonymous_functions83
// Make the tests pass!

// I AM NOT DONE
//
// iterator returns a literal handing out the next element on every call
// and false once the elements have run out.
// Practices an iterator closure with a position.
package main_test

import "testing"

func iterator(items []string) func() (string, bool) {
	pos := 0
	return func() (string, bool) {
		pos++
		return items[pos], true
	}
}

func TestIterator(t *testing.T) {
	next := iterator([]string{"a", "b"})
	var got []string
	for v, ok := next(); ok; v, ok = next() {
		got = append(got, v)
	}
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Errorf("iterated %v", got)
	}
}
