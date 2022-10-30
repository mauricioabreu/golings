package exercises

import (
	"fmt"
	"os/exec"
)

func Run(name string) (string, error) {
	exercise, err := Find(name)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("go", "run", fmt.Sprintf("./%s", exercise.Path))
	cOut, err := cmd.CombinedOutput()
	return string(cOut), err
}
