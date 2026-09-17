// generics100
// Make the tests pass!

// I AM NOT DONE
//
// SumAll adds up values of type Cents (type Cents int).
// The code does not compile: the constraint only accepts int.
// Without a ~ the named types are not in the type set.
package main_test

import "testing"

type Cents int

type Integer interface {
	int | int64
}

func SumAll[T Integer](vals ...T) T {
	var s T
	for _, v := range vals {
		s += v
	}
	return s
}

func TestSumAll(t *testing.T) {
	if SumAll(Cents(150), Cents(250)) != 400 {
		t.Errorf("SumAll = %d", SumAll(Cents(150), Cents(250)))
	}
}
