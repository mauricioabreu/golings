// slices4
// Make me compile!

// I AM NOT DONE
package main_test

import (
	"reflect"
	"testing"
)

func TestGetOnlyFirstName(t *testing.T) {
	names := []string{"John", "Maria", "Carl", "Peter"}
	firstName := names[3]

	if firstName != "John" {
		t.Errorf("firstName value must be John")
	}
}

func TestGetFirstTwoNames(t *testing.T) {
	names := []string{"John", "Maria", "Carl", "Peter"}
	firstTwoNames := names[5:10]
	expectedFirstTwoNames := []string{"John", "Maria"}

	if !reflect.DeepEqual(firstTwoNames, expectedFirstTwoNames) {
		t.Errorf("firstTwoNames should be %v, but got %v", expectedFirstTwoNames, firstTwoNames)
	}
}

func TestGetLastTwoNames(t *testing.T) {
	names := []string{"John", "Maria", "Carl", "Peter"}
	lastTwoNames := names[5:10]
	expectedLastTwoNames := []string{"Carl", "Peter"}

	if !reflect.DeepEqual(lastTwoNames, expectedLastTwoNames) {
		t.Errorf("lastTwoNames should be %v, but got %v", expectedLastTwoNames, lastTwoNames)
	}
}
