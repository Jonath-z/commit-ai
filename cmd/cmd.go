package cmd

import (
	"flag"
	"fmt"

	"github.com/Jonath-z/commit-ai/src"
)

var apiKey *string

func Cli() *string {
	name := flag.String("apiKey", "", "Openai API Key to use")
	flag.Parse()

	if *name == "" && apiKey == nil {
		fmt.Println("Please save the API key by using --apiKey <api key>")
	} else {
		gitDiff := src.GetGitChanges()
		apiKey = name
		commitMsg := src.GenerateCommitMessage(gitDiff, *apiKey)
		fmt.Println(commitMsg)
		fmt.Println("------------- Executing the commit message ------------")
		src.ExecuteCommitMsg(commitMsg)
	}

	return name
}
