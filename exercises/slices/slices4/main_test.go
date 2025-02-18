// slices4
// Make me compile!
package main_test

import (
	"reflect"
	"testing"
)

func TestGetOnlyFirstName(t *testing.T) {
	names := []string{"John", "Maria", "Carl", "Peter"}
	firstName := names[0]

	if firstName != "John" {
		t.Errorf("firstName value must be John")
	}
}

func TestGetFirstTwoNames(t *testing.T) {
	names := []string{"John", "Maria", "Carl", "Peter"}
	firstTwoNames := names[0:2]
	expectedFirstTwoNames := []string{"John", "Maria"}

	if !reflect.DeepEqual(firstTwoNames, expectedFirstTwoNames) {
		t.Errorf("firstTwoNames should be %v, but got %v", expectedFirstTwoNames, firstTwoNames)
	}
}

func TestGetLastTwoNames(t *testing.T) {
	names := []string{"John", "Maria", "Carl", "Peter"}
	lastTwoNames := names[2:4]
	expectedLastTwoNames := []string{"Carl", "Peter"}

	if !reflect.DeepEqual(lastTwoNames, expectedLastTwoNames) {
		t.Errorf("lastTwoNames should be %v, but got %v", expectedLastTwoNames, lastTwoNames)
	}
}
