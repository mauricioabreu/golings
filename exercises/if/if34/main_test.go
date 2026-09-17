// if34
// Make the tests pass!

// I AM NOT DONE
//
// season by month number: 12, 1 and 2 are "winter"; 3-5 "spring"; 6-8 "summer";
// 9-11 "autumn"; anything else "invalid".
// Practices range conditions with a special case.
package main_test

import "testing"

func season(month int) string {
	if month <= 2 {
		return "winter"
	} else if month <= 5 {
		return "spring"
	} else if month <= 8 {
		return "summer"
	}
	return "autumn"
}

func TestSeason(t *testing.T) {
	cases := map[int]string{1: "winter", 12: "winter", 3: "spring", 8: "summer", 11: "autumn", 0: "invalid", 13: "invalid"}
	for in, want := range cases {
		if got := season(in); got != want {
			t.Errorf("season(%d) = %s, want %s", in, got, want)
		}
	}
}
