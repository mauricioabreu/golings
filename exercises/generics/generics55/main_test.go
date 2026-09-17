// generics55
// Make the tests pass!

// I AM NOT DONE
//
// Result[T] holds a value or an error; Map applies a function only when there is no error.
// Practices generic wrapper types.
package main_test

import (
	"errors"
	"testing"
)

type Result[T any] struct {
	Val T
	Err error
}

func MapResult[T, U any](r Result[T], f func(T) U) Result[U] {
	return Result[U]{Val: f(r.Val)}
}

func TestMapResult(t *testing.T) {
	calls := 0
	double := func(x int) float64 { calls++; return float64(x) * 2 }
	ok := MapResult(Result[int]{Val: 4}, double)
	bad := MapResult(Result[int]{Err: errors.New("x")}, double)
	if ok.Val != 8 || ok.Err != nil || bad.Err == nil || calls != 1 {
		t.Errorf("ok=%+v bad=%+v calls=%d", ok, bad, calls)
	}
}
