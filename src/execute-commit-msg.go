package src

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommitMsg(commitMsg string, files []string) {
	addArgs := []string{"add"}
	if len(files) > 0 {
		addArgs = append(addArgs, "--")
		addArgs = append(addArgs, files...)
	} else {
		addArgs = append(addArgs, ".")
	}
	addCmd := exec.Command("git", addArgs...)
	addErr := addCmd.Run()

	if addErr != nil {
		fmt.Println("Could not stage the changes: ", addErr.Error())
		return
	}

	cmd := exec.Command("git", "commit", "-m", commitMsg)

	fmt.Println(cmd.Args)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer stdin.Close()
	buffer := new(bytes.Buffer)
	cmd.Stdout = buffer
	cmd.Stderr = os.Stderr

	if err = cmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
		return
	}

	cmd.Wait()
	exec.Command("git", "commit", "--amend")
}
