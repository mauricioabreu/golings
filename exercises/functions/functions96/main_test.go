// functions96
// Make the tests pass!

// I AM NOT DONE
//
// depositAll must top up every account in the map.
// The code does not compile: a pointer method cannot be called on a map element.
// Map elements are not addressable.
package main_test

import "testing"

type Account struct{ Balance int }

func (a *Account) Deposit(n int) { a.Balance += n }

func depositAll(accounts map[string]Account, n int) {
	for k := range accounts {
		accounts[k].Deposit(n)
	}
}

func TestDepositAll(t *testing.T) {
	accs := map[string]*Account{"ann": {Balance: 10}, "bob": {Balance: 0}}
	depositAll(accs, 5)
	if accs["ann"].Balance != 15 || accs["bob"].Balance != 5 {
		t.Errorf("balances = %d, %d; want 15, 5", accs["ann"].Balance, accs["bob"].Balance)
	}
}
