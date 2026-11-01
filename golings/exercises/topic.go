package exercises

import (
	"path/filepath"
	"strings"
)

// Topic is the folder an exercise lives in: for
// `exercises/variables/variables1/main.go` it is "variables". Exercises that
// aren't laid out as `exercises/<topic>/<name>/...` have no topic.
func (e Exercise) Topic() string {
	parts := strings.Split(filepath.ToSlash(e.Path), "/")
	for i, part := range parts {
		if part == "exercises" && len(parts) > i+3 {
			return parts[i+1]
		}
	}
	return ""
}

// SolutionPath is where the solution of the exercise lives, relative to the
// root of the solutions: `exercises/slices/slices1/main.go` becomes
// `solutions/slices/slices1/main.go`. It is empty for exercises that aren't
// laid out as `exercises/<topic>/<name>/...`.
func (e Exercise) SolutionPath() string {
	parts := strings.Split(filepath.ToSlash(e.Path), "/")
	for i, part := range parts {
		if part == "exercises" && len(parts) > i+3 {
			return "solutions/" + strings.Join(parts[i+1:], "/")
		}
	}
	return ""
}

type TopicProgress struct {
	Topic string
	Done  int
	Total int
}

// ProgressByTopic counts done and total exercises per topic, in the order
// topics first appear in the info file. Like Progress, it leaves out exercises
// the local Go can't run.
func ProgressByTopic(infoFile string) ([]TopicProgress, error) {
	allExercises, err := List(infoFile)
	if err != nil {
		return nil, err
	}

	var topics []TopicProgress
	index := map[string]int{}
	for _, exercise := range allExercises {
		if !runnable(exercise) {
			continue
		}
		topic := exercise.Topic()
		i, ok := index[topic]
		if !ok {
			i = len(topics)
			index[topic] = i
			topics = append(topics, TopicProgress{Topic: topic})
		}
		topics[i].Total++
		if exercise.State() == Done {
			topics[i].Done++
		}
	}
	return topics, nil
}
