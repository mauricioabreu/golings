// anonymous_functions89
// Make the tests pass!

// I AM NOT DONE
//
// Field holds a function validating a value.
// Practices function fields set from a literal.
package main_test

import "testing"

type Field struct {
	Name  string
	Check func(int) bool
}

var age = Field{
	Name:  "age",
	Check: func(v int) bool { return v >= 0 || v <= 150 },
}

func TestAgeCheck(t *testing.T) {
	if !age.Check(30) || age.Check(-1) || age.Check(200) {
		t.Errorf("age.Check works incorrectly")
	}
}
