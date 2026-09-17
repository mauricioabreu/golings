// switch48
// Make the tests pass!

// I AM NOT DONE
//
// priority for vehicles: "ambulance" and "fire" are 1; "police" is 2, or 1 with
// the siren on; everything else is 3.
// Practices a switch with an extra condition in a branch.
package main_test

import "testing"

func priority(vehicle string, siren bool) int {
	switch vehicle {
	case "ambulance":
		return 1
	case "police":
		return 2
	}
	return 3
}

func TestPriority(t *testing.T) {
	cases := []struct {
		v     string
		siren bool
		want  int
	}{{"ambulance", false, 1}, {"fire", false, 1}, {"police", false, 2}, {"police", true, 1}, {"taxi", true, 3}}
	for _, c := range cases {
		if got := priority(c.v, c.siren); got != c.want {
			t.Errorf("priority(%s, %v) = %d, want %d", c.v, c.siren, got, c.want)
		}
	}
}
