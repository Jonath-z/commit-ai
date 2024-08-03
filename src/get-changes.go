package src

import (
	"fmt"
	"os/exec"
)

func GetGitChanges() string {
	cmd := exec.Command("git", "diff")

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
