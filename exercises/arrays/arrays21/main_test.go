// arrays21
// Make the tests pass!

// I AM NOT DONE
//
// record updates a pupil's score in a map[string]Report, where Report holds
// a Scores array. The code does not compile: a map element cannot be changed in part.
// Map values are not addressable.
package main_test

import "testing"

type Report struct {
	Scores [3]int
}

func record(m map[string]Report, name string, i, score int) {
	m[name].Scores[i] = score
}

func TestRecord(t *testing.T) {
	m := map[string]Report{"ann": {}}
	record(m, "ann", 1, 5)
	record(m, "bob", 0, 4)
	if m["ann"].Scores != [3]int{0, 5, 0} || m["bob"].Scores != [3]int{4, 0, 0} {
		t.Errorf("m = %v", m)
	}
}
