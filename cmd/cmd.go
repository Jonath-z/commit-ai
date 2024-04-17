package cmd

import (
	"flag"
	"fmt"

	"github.com/Jonath-z/commit-ai/src"
)

func Cli() *string {
	name := flag.String("apiKey", "", "Openai API Key to use")
	flag.Parse()

	if *name == "" {
		fmt.Println("Please save the API key by using --apiKey <api key>")
	} else {
		gitDiff := src.GetGitChanges()
		commitMsg := src.GenerateCommitMessage(gitDiff, *name)
		fmt.Println(commitMsg)
		fmt.Println("------------- Executing the commit message ------------")
		src.ExecuteCommitMsg(commitMsg)
	}

	return name
}
