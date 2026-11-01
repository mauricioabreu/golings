package main

import (
	"runtime"
	"runtime/debug"
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

func TestResolveVersion(t *testing.T) {
	installed := &debug.BuildInfo{Main: debug.Module{Version: "v0.10.0"}}
	devel := &debug.BuildInfo{Main: debug.Module{Version: "(devel)"}}

	tests := []struct {
		name                              string
		version, commit, date             string
		info                              *debug.BuildInfo
		ok                                bool
		wantVersion, wantCommit, wantDate string
	}{
		{"go install of a tagged version", "dev", "none", "unknown", installed, true, "0.10.0", "", ""},
		{"stamped by goreleaser", "0.10.0", "abc", "2026-01-01", installed, true, "0.10.0", "abc", "2026-01-01"},
		{"built from a working copy", "dev", "none", "unknown", devel, true, "dev", "none", "unknown"},
		{"no build info", "dev", "none", "unknown", nil, false, "dev", "none", "unknown"},
		{"empty module version", "dev", "none", "unknown", &debug.BuildInfo{}, true, "dev", "none", "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			version, commit, date := resolveVersion(tt.version, tt.commit, tt.date, tt.info, tt.ok)

			if version != tt.wantVersion || commit != tt.wantCommit || date != tt.wantDate {
				t.Errorf("resolveVersion() = (%q, %q, %q), want (%q, %q, %q)", version, commit, date, tt.wantVersion, tt.wantCommit, tt.wantDate)
			}
		})
	}
}
