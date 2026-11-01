package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/cleonte/golings/golings/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	cmd.Execute(BuildVersion(resolveVersion(version, commit, date, info, ok)))
}

// resolveVersion fills in the version of builds that GoReleaser did not stamp.
// `go install module@v1.2.3` records the module version in the build info, so
// it is used when the version is still the "dev" default. Such builds have no
// commit or build date to report.
func resolveVersion(version, commit, date string, info *debug.BuildInfo, ok bool) (string, string, string) {
	if version != "dev" || !ok || info == nil {
		return version, commit, date
	}

	module := info.Main.Version
	if module == "" || module == "(devel)" {
		return version, commit, date
	}

	return strings.TrimPrefix(module, "v"), "", ""
}

func BuildVersion(version, commit, date string) string {
	result := version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	result = fmt.Sprintf("%s\ngoos: %s\ngoarch: %s", result, runtime.GOOS, runtime.GOARCH)

	return result
}
