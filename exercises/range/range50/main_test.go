// range50
// Make the tests pass!

// I AM NOT DONE
//
// byTeam groups the names of the players by team, keeping the input order inside a team.
// Practices a range over a slice of structs and a map of slices.
package main_test

import (
	"reflect"
	"testing"
)

type Player struct{ Name, Team string }

func byTeam(ps []Player) map[string][]string {
	out := map[string][]string{}
	for _, p := range ps {
		team := p.Name
		out[team] = []string{p.Name}
	}
	return out
}

func TestByTeam(t *testing.T) {
	got := byTeam([]Player{{"ann", "red"}, {"bob", "blue"}, {"cid", "red"}})
	if !reflect.DeepEqual(got, map[string][]string{"red": {"ann", "cid"}, "blue": {"bob"}}) {
		t.Errorf("byTeam = %v", got)
	}
}
