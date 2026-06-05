package providers

import (
	"fmt"

	"github.com/Jonath-z/commit-ai/src/config"
)

type Provider interface {
	Generate(systemPrompt, userPrompt string) (string, error)
}

func New(cfg config.Config) (Provider, error) {
	switch cfg.Provider {
	case "openai":
		if cfg.OpenAI.APIKey == "" {
			return nil, fmt.Errorf("openai api key not set. Run: commit-ai config set openai.api_key <key>")
		}
		return NewOpenAI(cfg.OpenAI.APIKey, cfg.OpenAI.Model, ""), nil
	case "local":
		if cfg.Local.BaseURL == "" {
			return nil, fmt.Errorf("local base_url not set. Run: commit-ai config set local.base_url <url>")
		}
		if cfg.Local.Model == "" {
			return nil, fmt.Errorf("local model not set. Run: commit-ai config set local.model <model>")
		}
		return NewOpenAI(cfg.Local.APIKey, cfg.Local.Model, cfg.Local.BaseURL), nil
	case "claude-code":
		return NewClaudeCode(cfg.ClaudeCode.Path, cfg.ClaudeCode.Model), nil
	case "codex":
		return NewCodex(cfg.Codex.Path, cfg.Codex.Model), nil
	case "":
		return nil, fmt.Errorf("no provider configured. Run: commit-ai config set provider <openai|claude-code|codex|local>")
	default:
		return nil, fmt.Errorf("unknown provider %q", cfg.Provider)
	}
}
