// anonymous_functions21
// Make the tests pass!

// I AM NOT DONE
//
// retry repeats an action with an exponential delay; the sleep function
// is passed in as a literal so the test does not wait.
// Practices injecting behaviour through a function parameter.
package main_test

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func retry(attempts int, base time.Duration, sleep func(time.Duration), f func() error) error {
	var err error
	delay := base
	for i := 0; i < attempts; i++ {
		if err = f(); err == nil {
			return nil
		}
		sleep(delay)
		delay += base
	}
	return err
}

func TestRetry(t *testing.T) {
	var slept []time.Duration
	err := retry(4, time.Second, func(d time.Duration) { slept = append(slept, d) }, func() error { return errors.New("down") })
	want := []time.Duration{time.Second, 2 * time.Second, 4 * time.Second}
	if err == nil || !reflect.DeepEqual(slept, want) {
		t.Errorf("slept = %v, err = %v", slept, err)
	}
}
