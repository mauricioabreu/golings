// switch93
// Make the tests pass!

// I AM NOT DONE
//
// action returns "go", "slow" and "stop" for the signals green, yellow and red.
// Practices a switch on a string.
package main_test

import "testing"

func action(light string) string {
	switch light {
	case "green":
		return "go"
	case "yellow":
		return "go"
	case "red":
		return "stop"
	}
	return "stop"
}

func TestAction(t *testing.T) {
	cases := map[string]string{"green": "go", "yellow": "slow", "red": "stop", "off": "stop"}
	for in, want := range cases {
		if got := action(in); got != want {
			t.Errorf("action(%s) = %s, want %s", in, got, want)
		}
	}
}
