// switch65
// Make the tests pass!

// I AM NOT DONE
//
// partOfDay returns "night" (0-5), "morning" (6-11), "day" (12-17), "evening" (18-23)
// and "invalid" for any other hour.
// Practices switch true with a validity check.
package main_test

import "testing"

func partOfDay(h int) string {
	switch {
	case h < 6:
		return "night"
	case h <= 12:
		return "morning"
	case h < 18:
		return "day"
	}
	return "evening"
}

func TestPartOfDay(t *testing.T) {
	cases := map[int]string{0: "night", 5: "night", 6: "morning", 12: "day", 18: "evening", 23: "evening", 24: "invalid", -1: "invalid"}
	for in, want := range cases {
		if got := partOfDay(in); got != want {
			t.Errorf("partOfDay(%d) = %s, want %s", in, got, want)
		}
	}
}
