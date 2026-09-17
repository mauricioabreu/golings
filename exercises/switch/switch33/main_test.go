// switch33
// Make the tests pass!

// I AM NOT DONE
//
// isWeekend must return true for Saturday and Sunday.
// Right now Saturday returns false.
// Go has no implicit fallthrough to the next case.
package main_test

import "testing"

func isWeekend(day string) bool {
	switch day {
	case "sat":
	case "sun":
		return true
	}
	return false
}

func TestIsWeekend(t *testing.T) {
	if !isWeekend("sat") || !isWeekend("sun") {
		t.Errorf("sat and sun are weekend days")
	}
	if isWeekend("mon") {
		t.Errorf("mon is not a weekend day")
	}
}
