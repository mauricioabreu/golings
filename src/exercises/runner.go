package exercises

import "os/exec"

func Run(exercise string) (string, error) {
	cmd := exec.Command("go", "run", "./exercises/"+exercise+".go")
	cOut, err := cmd.CombinedOutput()
	return string(cOut), err
}
