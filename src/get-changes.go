package src

import (
	"fmt"
	"os/exec"
	"strings"
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

	diff := string(output)

	for _, file := range untrackedFiles(files) {
		// --no-index exits with 1 when the files differ, so the error is expected
		untrackedDiff, _ := exec.Command("git", "diff", "--no-index", "--", "/dev/null", file).Output()
		diff += string(untrackedDiff)
	}

	if diff == " " {
		return "Initial commit"
	}
	return diff
}

func untrackedFiles(files []string) []string {
	args := []string{"ls-files", "--others", "--exclude-standard", "-z"}
	if len(files) > 0 {
		args = append(args, "--")
		args = append(args, files...)
	}
	output, err := exec.Command("git", args...).Output()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return strings.FieldsFunc(string(output), func(r rune) bool { return r == 0 })
}
