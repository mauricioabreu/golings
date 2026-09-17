// slices55
// Make the tests pass!

// I AM NOT DONE
//
// deleteAt removes an element from a slice of pointers. The objects removed have to become
// unreachable through the array, yet the last slot still refers to one of them.
// After the shift a stale pointer is left in the array past the length of the slice.
package main_test

import "testing"

type Conn struct{ ID int }

func deleteAt(s []*Conn, i int) []*Conn {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func TestDeleteAt(t *testing.T) {
	s := []*Conn{{1}, {2}, {3}}
	s = deleteAt(s, 0)
	if len(s) != 2 || s[0].ID != 2 || s[1].ID != 3 {
		t.Fatalf("deleteAt = %v", s)
	}
	if tail := s[:3][2]; tail != nil {
		t.Errorf("stale pointer left in backing array: %v", *tail)
	}
}
