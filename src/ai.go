package src

import (
	"github.com/Jonath-z/commit-ai/src/providers"
)

func prompt(s string) string {
	return `
	Given the git changes below, please draft a concise commit message that accurately summarizes the modifications. Follow these guidelines:

	1. Limit your commit message to 10 words.
	2. The whole commit message should in lowercase, no uppercase characters are allowed.
	3. Do not respond in Markdown

	   Git Changes:

		` + s
}

func GenerateCommitMessage(p providers.Provider, gitDiff string) (string, error) {
	return p.Generate(SystemPrompt, prompt(gitDiff))
}
