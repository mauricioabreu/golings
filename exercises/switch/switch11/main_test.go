// switch11
// Make the tests pass!

// I AM NOT DONE
//
// advance moves an order to its next state:
// new -> paid -> shipped -> done. done and unknown states give an error.
// Practices a switch as a transition table.
package main_test

import (
	"errors"
	"testing"
)

func advance(state string) (string, error) {
	switch state {
	case "new":
		return "paid", nil
	case "paid":
		return "done", nil
	case "shipped":
		return "done", nil
	}
	return state, nil
}

func TestAdvance(t *testing.T) {
	_ = errors.New
	chain := []string{"new", "paid", "shipped", "done"}
	for i := 0; i < len(chain)-1; i++ {
		if got, err := advance(chain[i]); err != nil || got != chain[i+1] {
			t.Errorf("advance(%s) = %s, %v; want %s", chain[i], got, err, chain[i+1])
		}
	}
	for _, s := range []string{"done", "lost"} {
		if _, err := advance(s); err == nil {
			t.Errorf("advance(%s) should fail", s)
		}
	}
}
