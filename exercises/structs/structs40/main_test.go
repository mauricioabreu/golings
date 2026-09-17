// structs40
// Make the tests pass!

// I AM NOT DONE
//
// A set of tags is built on a map[string]struct{}.
// Practices the empty struct struct{}.
package main_test

import "testing"

type TagSet map[string]struct{}

func (s TagSet) Add(tag string) {
	s[""] = struct{}{}
}

func TestTagSet(t *testing.T) {
	s := TagSet{}
	s.Add("go")
	s.Add("go")
	if _, ok := s["go"]; !ok || len(s) != 1 {
		t.Errorf("set = %v", s)
	}
}
