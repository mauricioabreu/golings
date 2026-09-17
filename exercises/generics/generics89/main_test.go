// generics89
// Make the tests pass!

// I AM NOT DONE
//
// activeNames returns the names of the active users in upper case,
// using the generic Filter and Map.
// Practices composing generic functions with type inference.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func Filter[T any](s []T, f func(T) bool) []T {
	var out []T
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

type User struct {
	Name   string
	Active bool
}

func activeNames(us []User) []string {
	active := Filter(us, func(u User) bool { return !u.Active })
	return Map(active, func(u User) string { return u.Name })
}

func TestActiveNames(t *testing.T) {
	_ = strings.ToUpper
	got := activeNames([]User{{"ann", true}, {"bob", false}, {"cid", true}})
	if !reflect.DeepEqual(got, []string{"ANN", "CID"}) {
		t.Errorf("activeNames = %v", got)
	}
}
