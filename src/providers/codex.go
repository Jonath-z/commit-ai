package providers

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type CodexProvider struct {
	path  string
	model string
}

func NewCodex(path, model string) *CodexProvider {
	if path == "" {
		path = "codex"
	}
	return &CodexProvider{path: path, model: model}
}

func (p *CodexProvider) Generate(systemPrompt, userPrompt string) (string, error) {
	args := []string{"exec"}
	if p.model != "" {
		args = append(args, "--model", p.model)
	}
	args = append(args, systemPrompt+"\n\n"+userPrompt)
	cmd := exec.Command(p.path, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("codex CLI failed: %w: %s", err, strings.TrimSpace(stderr.String()))
	}
	return strings.TrimSpace(stdout.String()), nil
}
