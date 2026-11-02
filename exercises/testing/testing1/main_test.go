// testing1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"slices"
	"testing"
)

type testCase struct {
	name string
	a, b int
	want int
}

var cases = []testCase{
	{name: "positive", a: 1, b: 2, want: 3},
	{name: "zero", a: 0, b: 0, want: 0},
	{name: "negative", a: -1, b: -2, want: -3},
}

func add(a, b int) int {
	return a + b
}

// runCases is the engine of a table-driven test: it calls check once for
// every case, each one as a subtest named after the case. When one fails,
// the report says which one.
func runCases(t *testing.T, cases []testCase, check func(t *testing.T, c testCase)) {
	for _, c := range cases {
		// TODO: Run each case as a subtest named after it.
		check(t, c)
	}
}

func TestAdd(t *testing.T) {
	runCases(t, cases, func(t *testing.T, c testCase) {
		if got := add(c.a, c.b); got != c.want {
			t.Errorf("add(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	})
}

func TestRunCasesNamesTheSubtests(t *testing.T) {
	var names []string

	runCases(t, cases, func(t *testing.T, c testCase) {
		names = append(names, t.Name())
	})

	want := []string{
		"TestRunCasesNamesTheSubtests/positive",
		"TestRunCasesNamesTheSubtests/zero",
		"TestRunCasesNamesTheSubtests/negative",
	}
	if !slices.Equal(names, want) {
		t.Errorf("the checks ran as %v, want %v", names, want)
	}
}
