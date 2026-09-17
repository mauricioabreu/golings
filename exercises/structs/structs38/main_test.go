// structs38
// Make the tests pass!

// I AM NOT DONE
//
// Mat2.Mul multiplies matrices and Mat2.Det computes the determinant.
// Practices a struct with an array field and methods on it.
package main_test

import "testing"

type Mat2 struct{ m [2][2]int }

func (a Mat2) Mul(b Mat2) Mat2 {
	var c Mat2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				c.m[i][j] += a.m[i][k] * b.m[j][k]
			}
		}
	}
	return c
}

func (a Mat2) Det() int {
	return a.m[0][0]*a.m[1][1] + a.m[0][1]*a.m[1][0]
}

func TestMat2(t *testing.T) {
	a := Mat2{[2][2]int{{1, 2}, {3, 4}}}
	b := Mat2{[2][2]int{{5, 6}, {7, 8}}}
	if got := a.Mul(b); got.m != [2][2]int{{19, 22}, {43, 50}} {
		t.Errorf("Mul = %v", got.m)
	}
	if a.Det() != -2 {
		t.Errorf("Det = %d", a.Det())
	}
}
