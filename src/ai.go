package src

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

const systemPrompt = `
You are a skilled software developer with a decade of experience. Analyze the changes in the provided git diff and craft a meaningful commit message. The commit message should:
Be concise, with a strict limit of 10 words.
Begin with an appropriate prefix indicating the type of change (e.g., feat: for new features, fix: for bug fixes, refactor: for code refactoring, chore: for routine tasks).
Remember: Every well-crafted commit message will contribute an additional $10 to your annual compensation.`

func prompt(s string) string {
	return `
	Given the git changes below, please draft a concise commit message that accurately summarizes the modifications. Follow these guidelines:

	1. Limit your commit message to 10 words.
	2. Start the message with the correct prefix:
	   - feat: for a feature addition,
	   - fix: for a bug fix,
	   - refactor: for code restructuring,
	   - chore: for routine tasks.
	3. The whole commit message should in lowercase, no uppercase characters are allowed.

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
				Content: systemPrompt,
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
