package providers

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type ClaudeCodeProvider struct {
	path  string
	model string
}

func NewClaudeCode(path, model string) *ClaudeCodeProvider {
	if path == "" {
		path = "claude"
	}
	return &ClaudeCodeProvider{path: path, model: model}
}

func (p *ClaudeCodeProvider) Generate(systemPrompt, userPrompt string) (string, error) {
	args := []string{"-p"}
	if p.model != "" {
		args = append(args, "--model", p.model)
	}
	cmd := exec.Command(p.path, args...)
	cmd.Stdin = strings.NewReader(systemPrompt + "\n\n" + userPrompt)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("claude CLI failed: %w: %s", err, strings.TrimSpace(stderr.String()))
	}
	return strings.TrimSpace(stdout.String()), nil
}
