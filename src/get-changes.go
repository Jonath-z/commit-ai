package src

import (
	"fmt"
	"os/exec"
)

func GetGitChanges(files []string) string {
	args := []string{"diff"}
	if len(files) > 0 {
		args = append(args, "--")
		args = append(args, files...)
	}
	cmd := exec.Command("git", args...)

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	cmd.Wait()
	if string(output) == " " {
		return "Initial commit"
	}
	return string(output)
}
