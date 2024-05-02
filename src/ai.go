package src

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func prompt(s string) string {
	return `
	Given the git changes below, please draft a concise commit message that accurately summarizes the modifications. Follow these guidelines:

	1. Limit your commit message to 10 words.
	2. The whole commit message should in lowercase, no uppercase characters are allowed.

	   Git Changes: 

		` + s
}

func GenerateCommitMessage(gitDiff string, apiKey string) string {
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: SystemPrompt,
			},
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
