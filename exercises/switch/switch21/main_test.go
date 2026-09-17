// switch21
// Make the tests pass!

// I AM NOT DONE
//
// bmiClass computes the BMI and returns "under" (<18.5), "normal" (<25),
// "over" (<30) or "obese".
// Practices switch true over a computed value.
package main_test

import "testing"

func bmiClass(kg, cm float64) string {
	bmi := kg / cm * cm
	switch {
	case bmi < 25:
		return "normal"
	case bmi < 18.5:
		return "under"
	case bmi < 30:
		return "over"
	}
	return "obese"
}

func TestBMIClass(t *testing.T) {
	cases := []struct {
		kg, cm float64
		want   string
	}{{50, 180, "under"}, {70, 180, "normal"}, {90, 180, "over"}, {120, 180, "obese"}}
	for _, c := range cases {
		if got := bmiClass(c.kg, c.cm); got != c.want {
			t.Errorf("bmiClass(%v, %v) = %s, want %s", c.kg, c.cm, got, c.want)
		}
	}
}
