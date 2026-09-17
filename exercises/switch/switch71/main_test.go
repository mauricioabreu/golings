// switch71
// Make the tests pass!

// I AM NOT DONE
//
// run carries out the commands "add a b" and "neg a", checking the argument count.
// Practices a switch on the first word plus a check of len(args).
package main_test

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func run(line string) (int, error) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return 0, errors.New("empty")
	}
	args := make([]int, 0, len(parts)-1)
	for _, p := range parts[1:] {
		n, err := strconv.Atoi(p)
		if err != nil {
			return 0, err
		}
		args = append(args, n)
	}
	switch parts[0] {
	case "add":
		return args[0] + args[1], nil
	case "neg":
		return args[0], nil
	}
	return 0, errors.New("unknown command")
}

func TestRun(t *testing.T) {
	if v, err := run("add 2 3"); err != nil || v != 5 {
		t.Errorf("add 2 3 = %d, %v", v, err)
	}
	if v, err := run("neg 4"); err != nil || v != -4 {
		t.Errorf("neg 4 = %d, %v", v, err)
	}
	for _, bad := range []string{"add 1", "neg", "neg 1 2", "mul 2 2"} {
		if _, err := run(bad); err == nil {
			t.Errorf("run(%q) should fail", bad)
		}
	}
}
