// if57
// Make the tests pass!

// I AM NOT DONE
//
// Thermostat turns the heating on below 18 degrees and off only above 22.
// Between the thresholds the state stays as it is.
// Practices conditions that depend on the current state.
package main_test

import "testing"

type Thermostat struct{ heating bool }

func (t *Thermostat) Update(temp float64) bool {
	if temp < 22 {
		t.heating = true
	} else {
		t.heating = false
	}
	return t.heating
}

func TestThermostat(t *testing.T) {
	th := &Thermostat{}
	temps := []float64{20, 17, 19, 21, 23, 20, 17.5}
	want := []bool{false, true, true, true, false, false, true}
	for i, temp := range temps {
		if got := th.Update(temp); got != want[i] {
			t.Errorf("step %d (%.1f°): heating = %v, want %v", i, temp, got, want[i])
		}
	}
}
