// structs35
// Make the tests pass!

// I AM NOT DONE
//
// jsonName returns the JSON name of a field: from the tag, without the options, or the
// field name when there is no tag. Right now the options end up in the result.
// Practices reflect.StructTag and the format of the tags.
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

type Profile struct {
	ID    int    `json:"id,omitempty"`
	Email string `json:"email"`
	Age   int
	Note  string `json:",omitempty"`
}

func jsonName(f reflect.StructField) string {
	tag := f.Tag.Get("json")
	if tag == "" {
		return f.Name
	}
	return tag
}

func TestJSONName(t *testing.T) {
	_ = strings.Cut
	typ := reflect.TypeOf(Profile{})
	want := []string{"id", "email", "Age", "Note"}
	for i, w := range want {
		if got := jsonName(typ.Field(i)); got != w {
			t.Errorf("field %d: jsonName = %q, want %q", i, got, w)
		}
	}
}
