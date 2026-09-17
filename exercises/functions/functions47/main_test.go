// functions47
// Make the tests pass!

// I AM NOT DONE
//
// factorial must return n! as a uint64, or ErrOverflow when it does not fit.
// Right now 21! quietly turns into a wrong number.
// Practices detecting overflow before the multiplication happens.
package main_test

import (
	"errors"
	"math"
	"testing"
)

var ErrOverflow = errors.New("overflow")

const limit = math.MaxUint64

func factorial(n int) (uint64, error) {
	var r uint64 = 1
	for i := 2; i <= n; i++ {
		r *= uint64(i)
	}
	return r, nil
}

func TestFactorial(t *testing.T) {
	if r, err := factorial(20); err != nil || r != 2432902008176640000 {
		t.Errorf("factorial(20) = %d, %v", r, err)
	}
	if _, err := factorial(21); !errors.Is(err, ErrOverflow) {
		t.Errorf("factorial(21) error = %v, want ErrOverflow", err)
	}
	if r, err := factorial(0); err != nil || r != 1 {
		t.Errorf("factorial(0) = %d, %v", r, err)
	}
}
