// if26
// Make the tests pass!

// I AM NOT DONE
//
// next returns the following signal: green -> yellow -> red -> green.
// An unknown signal gives "red", the safe value.
// Practices a chain of string comparisons.
package main_test

import "testing"

func next(light string) string {
	if light == "green" {
		return "yellow"
	} else if light == "yellow" {
		return "green"
	} else if light == "red" {
		return "yellow"
	}
	return "red"
}

func TestNext(t *testing.T) {
	cases := map[string]string{"green": "yellow", "yellow": "red", "red": "green", "blue": "red"}
	for in, want := range cases {
		if got := next(in); got != want {
			t.Errorf("next(%s) = %s, want %s", in, got, want)
		}
	}
}
