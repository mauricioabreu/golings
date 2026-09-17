// variables73
// Make the tests pass!

// I AM NOT DONE
//
// The package variables describe the address of a database.
// dsn must build a string such as "db.local:5432/shop".
// Practices declaring variables in a var ( ... ) group.
package main_test

import (
	"strconv"
	"testing"
)

var (
	dbHost = "db.local"
	dbPort = 3306
	dbName = "shop"
)

func dsn() string {
	return dbHost + ":" + dbName + "/" + strconv.Itoa(dbPort)
}

func TestDSN(t *testing.T) {
	if got := dsn(); got != "db.local:5432/shop" {
		t.Errorf("dsn() = %q, want %q", got, "db.local:5432/shop")
	}
}
