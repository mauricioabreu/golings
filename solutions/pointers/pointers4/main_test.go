// pointers4
// Make the tests pass!

package main_test

import "testing"

type Person struct {
	Name string
}

// greeting says hello to the person, or to nobody when there is no person.
func greeting(p *Person) string {
	if p == nil {
		return "Hello, nobody!"
	}
	return "Hello, " + p.Name + "!"
}

func TestGreeting(t *testing.T) {
	p := &Person{Name: "Ada"}

	if got := greeting(p); got != "Hello, Ada!" {
		t.Errorf("got %q, want %q", got, "Hello, Ada!")
	}
}

func TestGreetingWithoutAPerson(t *testing.T) {
	if got := greeting(nil); got != "Hello, nobody!" {
		t.Errorf("got %q, want %q", got, "Hello, nobody!")
	}
}
