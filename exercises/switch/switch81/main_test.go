// switch81
// Make the tests pass!

// I AM NOT DONE
//
// evalRPN evaluates an expression in reverse polish notation ("3 4 + 2 *").
// Practices a switch on a token and working with a stack.
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func evalRPN(expr string) int {
	var st []int
	for _, tok := range strings.Fields(expr) {
		switch tok {
		case "+", "-", "*":
			a, b := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, a+b)
		default:
			n, _ := strconv.Atoi(tok)
			st = append(st, n)
		}
	}
	return st[0]
}

func TestEvalRPN(t *testing.T) {
	cases := map[string]int{"3 4 + 2 *": 14, "10 3 -": 7, "2 3 4 * -": -10, "5": 5}
	for in, want := range cases {
		if got := evalRPN(in); got != want {
			t.Errorf("evalRPN(%q) = %d, want %d", in, got, want)
		}
	}
}
