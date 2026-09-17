// if102
// Make the tests pass!

// I AM NOT DONE
//
// bmiCategory: below 18.5 is "under", 25 and above is "over", the rest is "normal".
// Practices branching on fractional thresholds.
package main_test

import "testing"

func bmiCategory(bmi float64) string {
	if bmi < 18.5 {
		return "normal"
	}
	if bmi >= 25 {
		return "over"
	}
	return "normal"
}

func TestBMICategory(t *testing.T) {
	cases := map[float64]string{17: "under", 18.5: "normal", 24.9: "normal", 25: "over"}
	for in, want := range cases {
		if got := bmiCategory(in); got != want {
			t.Errorf("bmiCategory(%v) = %s, want %s", in, got, want)
		}
	}
}
