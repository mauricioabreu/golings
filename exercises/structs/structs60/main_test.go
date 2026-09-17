// structs60
// Make the tests pass!

// I AM NOT DONE
//
// Money.Add adds amounts of one currency; different currencies are an error.
// Practices methods returning (T, error).
package main_test

import (
	"errors"
	"testing"
)

type Money struct {
	Cents    int
	Currency string
}

func (m Money) Add(o Money) (Money, error) {
	m.Cents += o.Cents
	return m, nil
}

func TestMoneyAdd(t *testing.T) {
	_ = errors.New
	if s, err := (Money{150, "RUB"}).Add(Money{50, "RUB"}); err != nil || s != (Money{200, "RUB"}) {
		t.Errorf("Add = %v, %v", s, err)
	}
	if _, err := (Money{1, "RUB"}).Add(Money{1, "USD"}); err == nil {
		t.Errorf("adding different currencies should fail")
	}
}
