// structs25
// Make the tests pass!

// I AM NOT DONE
//
// Measurement structs with a "missing" value of NaN have to count as equal.
// == on structs compares float fields under IEEE rules, and NaN != NaN.
package main_test

import (
	"math"
	"testing"
)

type Measurement struct {
	Sensor string
	Value  float64
}

func (m Measurement) Equal(o Measurement) bool {
	return m == o
}

func TestMeasurementEqual(t *testing.T) {
	nan := math.NaN()
	if !(Measurement{"t1", nan}).Equal(Measurement{"t1", nan}) {
		t.Errorf("NaN measurements should be equal")
	}
	if (Measurement{"t1", 1}).Equal(Measurement{"t2", 1}) {
		t.Errorf("different sensors should differ")
	}
}
