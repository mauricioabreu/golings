// switch15
// Make the tests pass!

// I AM NOT DONE
//
// outcome decides a game of rock paper scissors with a switch
// over a struct holding the pair of moves.
// Practices comparable structs in a case.
package main_test

import "testing"

type pair struct{ a, b string }

func outcome(a, b string) string {
	if a == b {
		return "draw"
	}
	switch (pair{a, b}) {
	case pair{"rock", "paper"}, pair{"paper", "scissors"}:
		return "first"
	}
	return "draw"
}

func TestOutcome(t *testing.T) {
	cases := []struct{ a, b, want string }{
		{"rock", "scissors", "first"}, {"paper", "rock", "first"}, {"scissors", "paper", "first"},
		{"rock", "paper", "second"}, {"rock", "rock", "draw"},
	}
	for _, c := range cases {
		if got := outcome(c.a, c.b); got != c.want {
			t.Errorf("outcome(%s, %s) = %s, want %s", c.a, c.b, got, c.want)
		}
	}
}
