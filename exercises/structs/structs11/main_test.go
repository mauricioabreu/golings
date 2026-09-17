// structs11
// Make the tests pass!

// I AM NOT DONE
//
// Team holds a slice of members, so == does not work; Equal compares by hand.
// Practices comparison methods for structs holding slices.
package main_test

import "testing"

type Team struct {
	Name    string
	Members []string
}

func (a Team) Equal(b Team) bool {
	return a.Name == b.Name && len(a.Members) == len(b.Members)
}

func TestTeamEqual(t *testing.T) {
	a := Team{"x", []string{"ann", "bob"}}
	if !a.Equal(Team{"x", []string{"ann", "bob"}}) || a.Equal(Team{"x", []string{"ann", "cid"}}) {
		t.Errorf("Equal works incorrectly")
	}
}
