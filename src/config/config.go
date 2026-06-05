package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type OpenAIConfig struct {
	APIKey string `json:"api_key"`
	Model  string `json:"model"`
}

type LocalConfig struct {
	BaseURL string `json:"base_url"`
	Model   string `json:"model"`
	APIKey  string `json:"api_key,omitempty"`
}

type ClaudeCodeConfig struct {
	Path  string `json:"path,omitempty"`
	Model string `json:"model,omitempty"`
}

type CodexConfig struct {
	Path  string `json:"path,omitempty"`
	Model string `json:"model,omitempty"`
}

type Config struct {
	Provider   string           `json:"provider"`
	OpenAI     OpenAIConfig     `json:"openai"`
	Local      LocalConfig      `json:"local"`
	ClaudeCode ClaudeCodeConfig `json:"claude_code"`
	Codex      CodexConfig      `json:"codex"`
}

func defaultConfig() Config {
	return Config{
		Provider: "openai",
		OpenAI:   OpenAIConfig{Model: "gpt-4o"},
		Local: LocalConfig{
			BaseURL: "http://localhost:11434/v1",
			Model:   "llama3",
		},
	}
}

func Path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "commit-ai", "config.json"), nil
}

func Load() (Config, error) {
	cfg := defaultConfig()
	p, err := Path()
	if err != nil {
		return cfg, err
	}
	data, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("parsing %s: %w", p, err)
	}
	return cfg, nil
}

func Save(cfg Config) error {
	p, err := Path()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0o600)
}

func ValidProvider(name string) bool {
	switch name {
	case "openai", "claude-code", "codex", "local":
		return true
	}
	return false
}

func Set(cfg *Config, key, value string) error {
	switch key {
	case "provider":
		if !ValidProvider(value) {
			return fmt.Errorf("unknown provider %q (valid: openai, claude-code, codex, local)", value)
		}
		cfg.Provider = value
	case "openai.api_key":
		cfg.OpenAI.APIKey = value
	case "openai.model":
		cfg.OpenAI.Model = value
	case "local.base_url":
		cfg.Local.BaseURL = value
	case "local.model":
		cfg.Local.Model = value
	case "local.api_key":
		cfg.Local.APIKey = value
	case "claude_code.path":
		cfg.ClaudeCode.Path = value
	case "claude_code.model":
		cfg.ClaudeCode.Model = value
	case "codex.path":
		cfg.Codex.Path = value
	case "codex.model":
		cfg.Codex.Model = value
	default:
		return fmt.Errorf("unknown config key %q", key)
	}
	return nil
}
