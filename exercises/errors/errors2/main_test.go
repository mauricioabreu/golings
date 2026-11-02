// errors2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"testing"
)

// Sentinel errors: callers compare against these exact values.
var (
	ErrNotFound        = errors.New("item not found")
	ErrInvalidQuantity = errors.New("quantity must be positive")
)

var prices = map[string]int{"apple": 3, "pear": 5}

// total returns what the quantity of an item costs.
func total(item string, quantity int) (int, error) {
	// TODO: Return the sentinel errors above, not new errors that merely look the same.
	if quantity <= 0 {
		return 0, errors.New("quantity must be positive")
	}
	price, ok := prices[item]
	if !ok {
		return 0, errors.New("item not found")
	}
	return price * quantity, nil
}

func TestTotal(t *testing.T) {
	got, err := total("pear", 2)

	if err != nil || got != 10 {
		t.Errorf("total(pear, 2) = (%d, %v), want (10, nil)", got, err)
	}
}

func TestTotalOfAnUnknownItem(t *testing.T) {
	_, err := total("kiwi", 1)

	if err != ErrNotFound {
		t.Errorf("total(kiwi, 1) returned %v, want ErrNotFound", err)
	}
}

func TestTotalWithAnInvalidQuantity(t *testing.T) {
	_, err := total("apple", 0)

	if err != ErrInvalidQuantity {
		t.Errorf("total(apple, 0) returned %v, want ErrInvalidQuantity", err)
	}
}
