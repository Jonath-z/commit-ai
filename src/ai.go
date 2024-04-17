package src

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func prompt(s string) string {
	return `User You are a software developer with 10years of experience in the industry, found the changes on this git diff context, and make a meaning full commit. 
	Note: 
		- your commit sould have 10 words
	    - Foreach good commit 10 dollars is added the annual compensasion
	    - Add the prefix, prefix should be feat: for a feature, fix: for a fix, refactor: for a refactor, chore: for a chore change ` + s
}

func GenerateCommitMessage(gitDiff string, apiKey string) string {
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: prompt(gitDiff),
			},
		},
	})

	if err != nil {
		fmt.Print(err.Error())
	}

	return resp.Choices[0].Message.Content
}
