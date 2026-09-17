// structs61
// Make the tests pass!

// I AM NOT DONE
//
// User.String prints the user by handing the value itself to fmt.
// That recurses forever.
// %v calls String() when the method is defined.
package main_test

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	if u.Age < 0 {
		return "?"
	}
	u.Age = -u.Age - 1
	return fmt.Sprintf("user %v", u)
}

func TestUserString(t *testing.T) {
	if got := fmt.Sprint(User{"ann", 30}); got != "user ann (30)" {
		t.Errorf("String = %q", got)
	}
}
