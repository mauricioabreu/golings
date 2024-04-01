package exercises

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var ErrExerciseNotFound = errors.New("exercise not found")
var ErrNoPendingExercises = errors.New("no pending exercises")

type Info struct {
	Exercises []Exercise
}

func List(infoFile string) ([]Exercise, error) {
	var info Info

	data, err := os.ReadFile(infoFile)
	if err != nil {
		return info.Exercises, err
	}

	if err := toml.Unmarshal(data, &info); err != nil {
		return info.Exercises, err
	}

	return info.Exercises, nil
}

func NextPending(infoFile string) (Exercise, error) {
	allExercises, err := List(infoFile)
	if err != nil {
		return Exercise{}, err
	}

	for _, exercise := range allExercises {
		if exercise.State() == Pending {
			return exercise, nil
		}
	}

	return Exercise{}, ErrNoPendingExercises
}

func Find(exercise string, infoFile string) (Exercise, error) {
	exs, err := List(infoFile)
	if err != nil {
		return Exercise{}, err
	}

	for _, ex := range exs {
		if ex.Name == exercise {
			return ex, nil
		}
	}

	return Exercise{}, ErrExerciseNotFound
}

func Progress(infoFile string) (float32, int, int, error) {
	allExercises, err := List(infoFile)
	if err != nil {
		return 0.0, 0, 0, err
	}
	done := []Exercise{}
	for _, exercise := range allExercises {
		if exercise.State() == Done {
			done = append(done, exercise)
		}
	}

	totalDone := len(done)
	total := len(allExercises)

	return float32(totalDone) / float32(total), totalDone, total, nil
}
