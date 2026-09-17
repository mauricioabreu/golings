// structs62
// Make the tests pass!

// I AM NOT DONE
//
// Inventory.Remove lowers the quantity and drops the item at zero;
// more than there is cannot be taken away.
// Practices methods with checks over a map field.
package main_test

import (
	"errors"
	"testing"
)

type Inventory struct{ stock map[string]int }

func (inv *Inventory) Remove(item string, n int) error {
	inv.stock[item] -= n
	return nil
}

func TestInventory(t *testing.T) {
	_ = errors.New
	inv := &Inventory{stock: map[string]int{"pen": 3}}
	if err := inv.Remove("pen", 5); err == nil {
		t.Errorf("removing too much should fail")
	}
	inv.Remove("pen", 1)
	inv.Remove("pen", 2)
	if _, ok := inv.stock["pen"]; ok {
		t.Errorf("pen should be removed: %v", inv.stock)
	}
}
