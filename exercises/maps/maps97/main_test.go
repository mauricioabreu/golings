// maps97
// Make the tests pass!

// I AM NOT DONE
//
// Sparse keeps only the non-zero elements in a map[[2]int]int.
// Writing a zero has to remove the element.
// Practices delete for keeping a structure sparse.
package main_test

import "testing"

type Sparse map[[2]int]int

func (s Sparse) Set(r, c, v int) {
	s[[2]int{c, r}] = v
}

func (s Sparse) Get(r, c int) int { return s[[2]int{r, c}] }

func TestSparse(t *testing.T) {
	s := Sparse{}
	s.Set(1, 2, 5)
	s.Set(3, 3, 7)
	s.Set(3, 3, 0)
	if s.Get(1, 2) != 5 || s.Get(2, 1) != 0 || len(s) != 1 {
		t.Errorf("sparse = %v", s)
	}
}
