// switch74
// Make the tests pass!

// I AM NOT DONE
//
// canMove checks a task moving between states:
// todo -> doing, doing -> done or todo, done -> nowhere.
// Practices a switch on the source state with a check of the target.
package main_test

import "testing"

func canMove(from, to string) bool {
	switch from {
	case "todo":
		return to == "doing"
	case "doing":
		return to == "done"
	case "done":
		return true
	}
	return false
}

func TestCanMove(t *testing.T) {
	cases := []struct {
		from, to string
		want     bool
	}{{"todo", "doing", true}, {"todo", "done", false}, {"doing", "done", true}, {"doing", "todo", true}, {"done", "todo", false}, {"x", "todo", false}}
	for _, c := range cases {
		if got := canMove(c.from, c.to); got != c.want {
			t.Errorf("canMove(%s, %s) = %v, want %v", c.from, c.to, got, c.want)
		}
	}
}
