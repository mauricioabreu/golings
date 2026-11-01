package main

import (
	"runtime"
	"strings"
	"testing"
)

func TestBuildVersion(t *testing.T) {
	got := BuildVersion("v1.2.3", "abc123", "2026-01-01")

	for _, want := range []string{"v1.2.3", "commit: abc123", "built at: 2026-01-01", "goos: " + runtime.GOOS, "goarch: " + runtime.GOARCH} {
		if !strings.Contains(got, want) {
			t.Errorf("BuildVersion() = %q, missing %q", got, want)
		}
	}
}

func TestBuildVersionOmitsEmptyFields(t *testing.T) {
	got := BuildVersion("dev", "", "")

	for _, unwanted := range []string{"commit:", "built at:"} {
		if strings.Contains(got, unwanted) {
			t.Errorf("BuildVersion() = %q, should not contain %q", got, unwanted)
		}
	}
}
