// structs9
// Make the tests pass!

// I AM NOT DONE
//
// Dog embeds Animal, and has to say "woof" while using the name from Animal.
// A method on the outer type hides the promoted one.
package main_test

import "testing"

type Animal struct{ Name string }

func (a Animal) Sound() string { return a.Name + ": ..." }

type Dog struct{ Animal }

func (d Dog) Speak() string { return d.Name + ": woof" }
func (d Dog) Bark() string  { return "woof" }

type Sounder interface{ Sound() string }

func TestDogSound(t *testing.T) {
	var s Sounder = Dog{Animal{"Rex"}}
	if s.Sound() != "Rex: woof" {
		t.Errorf("Sound = %q", s.Sound())
	}
	if (Animal{"Cat"}).Sound() != "Cat: ..." {
		t.Errorf("Animal.Sound changed")
	}
}
