// arrays14
// Make the tests pass!

// I AM NOT DONE
//
// snapshot must return an independent copy of the readings that later measurements
// cannot touch. Right now the snapshot changes along with the array.
// The slice arr[:] refers to the same array.
package main_test

import "testing"

type Sensor struct {
	readings [4]int
}

func (s *Sensor) snapshot() []int {
	return s.readings[:]
}

func TestSnapshot(t *testing.T) {
	s := &Sensor{readings: [4]int{1, 2, 3, 4}}
	snap := s.snapshot()
	s.readings[0] = 100
	if snap[0] != 1 {
		t.Errorf("snapshot changed: %v", snap)
	}
}
