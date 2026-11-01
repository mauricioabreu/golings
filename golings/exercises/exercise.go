package exercises

import (
	"os"
	"regexp"
)

var notDoneRegex = regexp.MustCompile(`(?m)^\s*///?\s*I\s+AM\s+NOT\s+DONE`)

type Exercise struct {
	Name string
	Path string
	Mode string
	Hint string
	// GoVersion is the minimum Go version (e.g. "1.26") needed to run the exercise.
	GoVersion string `toml:"go_version"`
}

func (e Exercise) State() State {
	data, err := os.ReadFile(e.Path)
	if err != nil {
		return Pending
	}

	if notDoneRegex.Match(data) {
		return Pending
	}
	return Done
}

type State int

const (
	Pending State = iota + 1
	Done
)

func (s State) String() string {
	return [...]string{"Pending", "Done"}[s-1]
}
