// maps41
// Make the tests pass!

// I AM NOT DONE
//
// byCourse turns a map from pupil to courses into one from course to pupils, sorted.
// Practices inverting a many to many relation.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func byCourse(enroll map[string][]string) map[string][]string {
	out := map[string][]string{}
	for student, courses := range enroll {
		for _, c := range courses {
			out[student] = append(out[student], c)
		}
	}
	return out
}

func TestByCourse(t *testing.T) {
	got := byCourse(map[string][]string{"ann": {"go", "db"}, "bob": {"go"}})
	if !reflect.DeepEqual(got, map[string][]string{"go": {"ann", "bob"}, "db": {"ann"}}) {
		t.Errorf("byCourse = %v", got)
	}
}
