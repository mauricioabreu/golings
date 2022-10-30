package exercises

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type State int

const (
	Pending State = iota + 1
	Done
)

func (s State) String() string {
	return [...]string{"Pending", "Done"}[s-1]
}

type Exercise struct {
	Name  string
	Path  string
	Mode  string
	Hint  string
	State State
}

type Info struct {
	Exercises []Exercise
}

func List() ([]Exercise, error) {
	var info Info

	data, err := os.ReadFile("info.toml")
	if err != nil {
		return info.Exercises, err
	}

	if err := toml.Unmarshal(data, &info); err != nil {
		return info.Exercises, err
	}

	return info.Exercises, nil
}
