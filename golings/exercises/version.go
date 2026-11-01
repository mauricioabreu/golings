package exercises

import (
	"errors"
	"go/version"
	"os/exec"
	"strings"
	"sync"
)

var ErrGoVersionTooOld = errors.New("installed Go is too old for this exercise")

// LocalGoVersion reports the version of the `go` toolchain exercises are run with,
// e.g. "go1.26.0". It is looked up once per process.
var LocalGoVersion = sync.OnceValues(func() (string, error) {
	out, err := exec.Command("go", "env", "GOVERSION").Output()
	if err != nil {
		return "", err
	}
	fields := strings.Fields(string(out))
	if len(fields) == 0 {
		return "", errors.New("empty GOVERSION")
	}
	return fields[0], nil
})

// Supports reports whether the given local Go version (e.g. "go1.26.0") meets
// the exercise's minimum. Exercises without a go_version always match.
func (e Exercise) Supports(local string) bool {
	if e.GoVersion == "" {
		return true
	}
	return version.Compare(local, "go"+e.GoVersion) >= 0
}
