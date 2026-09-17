// concurrent95
// Make the tests pass!

// I AM NOT DONE
//
// A single sync.Once is used both for loading the config and for opening
// the database.
// The second function never runs.
// Practices that Once runs only the first function given to it.
package main_test

import (
	"sync"
	"testing"
)

type App struct {
	once   sync.Once
	config string
	db     string
}

func (a *App) Config() string {
	a.once.Do(func() { a.config = "loaded" })
	return a.config
}

func (a *App) DB() string {
	a.once.Do(func() { a.db = "connected" })
	return a.db
}

func TestApp(t *testing.T) {
	var a App
	if a.Config() != "loaded" || a.DB() != "connected" {
		t.Errorf("config = %q, db = %q", a.config, a.db)
	}
}
