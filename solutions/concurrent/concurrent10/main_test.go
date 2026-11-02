// concurrent10
// Make the tests pass!

package main_test

import (
	"sync"
	"testing"
)

// Settings loads its value the first time somebody asks for it, and every
// later call gets the same value, without loading it again.
type Settings struct {
	once  sync.Once
	value string
	loads int
}

func (s *Settings) Get() string {
	s.once.Do(func() {
		s.value = load()
		s.loads++
	})
	return s.value
}

func load() string {
	return "production"
}

func TestSettingsLoadOnlyOnce(t *testing.T) {
	var settings Settings
	var wg sync.WaitGroup

	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got := settings.Get(); got != "production" {
				t.Errorf("Get() = %q, want production", got)
			}
		}()
	}
	wg.Wait()

	if settings.loads != 1 {
		t.Errorf("the settings were loaded %d times, want once", settings.loads)
	}
}
