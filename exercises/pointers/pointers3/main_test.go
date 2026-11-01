// pointers3
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

type Person struct {
	Name string
	Age  int
}

// newPerson creates a person and returns a pointer to it.
func newPerson(name string, age int) *Person {
	// TODO: Return a pointer to a new person.
	return nil
}

// haveBirthday makes the person one year older.
func haveBirthday(p *Person) {
	// TODO: Make the person one year older.
}

func TestNewPerson(t *testing.T) {
	p := newPerson("Ada", 36)

	if p == nil {
		t.Fatal("newPerson should return a pointer to a Person, but got nil")
	}
	if p.Name != "Ada" || p.Age != 36 {
		t.Errorf("got %+v, want {Name:Ada Age:36}", *p)
	}
}

func TestHaveBirthday(t *testing.T) {
	p := newPerson("Ada", 36)
	if p == nil {
		t.Fatal("newPerson should return a pointer to a Person, but got nil")
	}

	haveBirthday(p)

	if p.Age != 37 {
		t.Errorf("Ada should be 37 after her birthday, but is %d", p.Age)
	}
}
