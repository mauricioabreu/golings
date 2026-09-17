// structs19
// Make the tests pass!

// I AM NOT DONE
//
// Student.Average returns the average mark, 0 with no marks, and Best the highest one.
// Practices methods over a slice field.
package main_test

import "testing"

type Student struct{ Grades []int }

func (s Student) Average() float64 {
	sum := 0
	for _, g := range s.Grades {
		sum += g
	}
	return float64(sum / len(s.Grades))
}

func (s Student) Best() int {
	best := 0
	for _, g := range s.Grades {
		if g > best {
			best = g
		}
	}
	return best
}

func TestStudent(t *testing.T) {
	s := Student{[]int{4, 5, 5, 3}}
	if s.Average() != 4.25 || s.Best() != 5 {
		t.Errorf("Average = %v, Best = %d", s.Average(), s.Best())
	}
	if (Student{}).Average() != 0 {
		t.Errorf("empty Average should be 0")
	}
}
