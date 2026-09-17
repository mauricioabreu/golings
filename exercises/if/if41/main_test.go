// if41
// Make the tests pass!

// I AM NOT DONE
//
// nextTicket hands out a number from the queue when the queue is not empty.
// pop is called both in the condition and in the body, so numbers go missing.
// The expressions in an if condition are evaluated on every check.
package main_test

import (
	"reflect"
	"testing"
)

type Queue struct{ items []int }

func (q *Queue) pop() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}

func nextTicket(q *Queue) int {
	if _, ok := q.pop(); ok {
		v, _ := q.pop()
		return v
	}
	return -1
}

func TestNextTicket(t *testing.T) {
	q := &Queue{items: []int{1, 2, 3}}
	got := []int{nextTicket(q), nextTicket(q), nextTicket(q), nextTicket(q)}
	if want := []int{1, 2, 3, -1}; !reflect.DeepEqual(got, want) {
		t.Errorf("tickets = %v, want %v", got, want)
	}
}
