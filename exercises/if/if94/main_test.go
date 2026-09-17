// if94
// Make the tests pass!

// I AM NOT DONE
//
// ageGroup returns "child" (<13), "teen" (13-17) or "adult" (18+),
// and an error for a negative age or an age above 150.
// Practices validating the input before the main logic.
package main_test

import (
	"errors"
	"testing"
)

func ageGroup(age int) (string, error) {
	if age < 13 {
		return "child", nil
	} else if age <= 18 {
		return "teen", nil
	}
	return "adult", nil
}

func TestAgeGroup(t *testing.T) {
	_ = errors.New
	cases := map[int]string{0: "child", 12: "child", 13: "teen", 17: "teen", 18: "adult"}
	for in, want := range cases {
		if got, err := ageGroup(in); err != nil || got != want {
			t.Errorf("ageGroup(%d) = %s, %v; want %s", in, got, err, want)
		}
	}
	for _, bad := range []int{-1, 151} {
		if _, err := ageGroup(bad); err == nil {
			t.Errorf("ageGroup(%d) should fail", bad)
		}
	}
}
