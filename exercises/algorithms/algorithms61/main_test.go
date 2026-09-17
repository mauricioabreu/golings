// algorithms61
// Make the tests pass!

// I AM NOT DONE
//
// Pattern: Eulerian path (Hierholzer's algorithm). Given a list of [from, to]
// tickets, build an itinerary starting at "JFK" that uses all the tickets.
// When choosing, take the lexicographically smaller airport.
// Expected asymptotics: O(E·log E) time, O(E) space.
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func findItinerary(tickets [][2]string) []string {
	return nil
}

func TestFindItinerary(t *testing.T) {
	_ = sort.Strings
	got := findItinerary([][2]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}})
	if !reflect.DeepEqual(got, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}) {
		t.Errorf("findItinerary = %v", got)
	}
	got = findItinerary([][2]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}})
	if !reflect.DeepEqual(got, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}) {
		t.Errorf("findItinerary = %v", got)
	}
	if got := findItinerary(nil); !reflect.DeepEqual(got, []string{"JFK"}) {
		t.Errorf("no tickets = %v", got)
	}
}
