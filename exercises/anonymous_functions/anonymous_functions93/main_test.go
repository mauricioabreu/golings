// anonymous_functions93
// Make the tests pass!

// I AM NOT DONE
//
// exec runs a command from a table of literals with its arguments;
// an unknown command or a wrong number of arguments is an error.
// Practices map[string]func([]int) (int, error).
package main_test

import (
	"errors"
	"testing"
)

var commands = map[string]func([]int) (int, error){
	"sum": func(a []int) (int, error) {
		s := 0
		for _, v := range a {
			s += v
		}
		return s, nil
	},
	"neg": func(a []int) (int, error) {
		return -a[0], nil
	},
}

func exec(name string, args ...int) (int, error) {
	return commands[name](args)
}

func TestExec(t *testing.T) {
	if v, err := exec("sum", 1, 2, 3); err != nil || v != 6 {
		t.Errorf("sum = %d, %v", v, err)
	}
	if _, err := exec("neg"); err == nil {
		t.Errorf("neg without args should fail")
	}
	if _, err := exec("mul", 2); err == nil {
		t.Errorf("unknown command should fail")
	}
}
