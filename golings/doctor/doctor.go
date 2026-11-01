// Package doctor diagnoses the learner's setup: Go toolchain, workspace and
// the optional requirements of individual exercises.
package doctor

import (
	"fmt"
	"go/version"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/cleonte/golings/golings/exercises"
)

// MinGoVersion is the oldest Go release golings supports.
const MinGoVersion = "1.26"

type Status int

const (
	OK Status = iota
	Warn
	Fail
)

type Check struct {
	Name   string
	Status Status
	Detail string
}

// Env holds everything doctor learns about the machine, so tests can fake it.
type Env struct {
	GoVersion    func() (string, error)
	CGOEnabled   func() bool
	HasCCompiler func() bool
	FileExists   func(path string) bool
}

func DefaultEnv() Env {
	return Env{
		GoVersion: exercises.LocalGoVersion,
		CGOEnabled: func() bool {
			out, err := exec.Command("go", "env", "CGO_ENABLED").Output()
			return err == nil && strings.TrimSpace(string(out)) == "1"
		},
		HasCCompiler: func() bool {
			for _, cc := range []string{"cc", "gcc", "clang"} {
				if _, err := exec.LookPath(cc); err == nil {
					return true
				}
			}
			return false
		},
		FileExists: func(path string) bool {
			_, err := os.Stat(path)
			return err == nil
		},
	}
}

// Run executes all checks and returns their results in display order.
func Run(infoFile string, env Env) []Check {
	checks := []Check{}

	local, err := env.GoVersion()
	checks = append(checks, goCheck(local, err))

	all, err := exercises.List(infoFile)
	if err != nil {
		return append(checks, Check{
			Name:   "Exercise manifest",
			Status: Fail,
			Detail: fmt.Sprintf("can't read %s: %v (run golings from the folder that holds it, or create one with `golings init`)", infoFile, err),
		})
	}
	checks = append(checks, Check{Name: "Exercise manifest", Status: OK, Detail: fmt.Sprintf("%s lists %d exercises", infoFile, len(all))})

	checks = append(checks, filesCheck(all, env))
	if local != "" {
		checks = append(checks, versionGatedCheck(all, local))
	}
	checks = append(checks, raceCheck(all, env))

	return checks
}

// Failed reports whether any check failed (warnings don't count).
func Failed(checks []Check) bool {
	for _, c := range checks {
		if c.Status == Fail {
			return true
		}
	}
	return false
}

func goCheck(local string, err error) Check {
	c := Check{Name: "Go toolchain"}
	switch {
	case err != nil:
		c.Status, c.Detail = Fail, fmt.Sprintf("can't run `go` (%v): is Go installed and on your PATH?", err)
	case version.Compare(local, "go"+MinGoVersion) < 0:
		c.Status, c.Detail = Fail, fmt.Sprintf("%s is too old, golings needs Go %s or newer", local, MinGoVersion)
	default:
		c.Status, c.Detail = OK, fmt.Sprintf("%s on %s/%s", local, runtime.GOOS, runtime.GOARCH)
	}
	return c
}

func filesCheck(all []exercises.Exercise, env Env) Check {
	var missing []string
	for _, ex := range all {
		if !env.FileExists(ex.Path) {
			missing = append(missing, ex.Name)
		}
	}
	if len(missing) > 0 {
		return Check{Name: "Exercise files", Status: Fail, Detail: fmt.Sprintf("%d missing: %s", len(missing), strings.Join(missing, ", "))}
	}
	return Check{Name: "Exercise files", Status: OK, Detail: "all present"}
}

func versionGatedCheck(all []exercises.Exercise, local string) Check {
	var tooNew []string
	for _, ex := range all {
		if !ex.Supports(local) {
			tooNew = append(tooNew, fmt.Sprintf("%s (needs Go %s)", ex.Name, ex.GoVersion))
		}
	}
	if len(tooNew) > 0 {
		return Check{Name: "Exercise Go versions", Status: Warn, Detail: fmt.Sprintf("%d will be skipped: %s", len(tooNew), strings.Join(tooNew, ", "))}
	}
	return Check{Name: "Exercise Go versions", Status: OK, Detail: "every exercise can run on this Go"}
}

// raceCheck looks at what `go test -race`, used by test-mode exercises, needs.
func raceCheck(all []exercises.Exercise, env Env) Check {
	c := Check{Name: "Race detector"}
	hasTests := false
	for _, ex := range all {
		if ex.Mode == "test" {
			hasTests = true
			break
		}
	}
	switch {
	case !hasTests:
		c.Status, c.Detail = OK, "not needed"
	case !env.CGOEnabled():
		c.Status, c.Detail = Warn, "CGO is disabled; test exercises run with -race and need CGO_ENABLED=1"
	case !env.HasCCompiler():
		c.Status, c.Detail = Warn, "no C compiler found (cc, gcc, clang); test exercises run with -race and need one"
	default:
		c.Status, c.Detail = OK, "CGO enabled and a C compiler is available"
	}
	return c
}
