// generics46
// Make the tests pass!

// I AM NOT DONE
//
// Scale multiplies a value by 1000. The code does not compile: the constraint holds int8,
// which cannot hold 1000.
// Constants in generic code have to fit every type of the set.
package main_test

import "testing"

type Num interface {
	~int8 | ~int | ~int64
}

func Scale[T Num](v T) T {
	return v * 1000
}

func TestScale(t *testing.T) {
	if Scale(3) != 3000 || Scale(int64(2)) != 2000 {
		t.Errorf("Scale works incorrectly")
	}
}
