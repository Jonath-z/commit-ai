package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Jonath-z/commit-ai/src/config"
)

func runConfig(args []string) {
	if len(args) == 0 {
		printConfigUsage(os.Stderr)
		os.Exit(1)
	}
	if isHelpArg(args[0]) {
		printConfigUsage(os.Stdout)
		return
	}
	switch args[0] {
	case "set":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: commit-ai config set <key> <value>")
			os.Exit(1)
		}
		cfg, err := config.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
			os.Exit(1)
		}
		if err := config.Set(&cfg, args[1], args[2]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := config.Save(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "failed to save config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Set %s\n", args[1])
	case "show":
		cfg, err := config.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
			os.Exit(1)
		}
		redacted := cfg
		redacted.OpenAI.APIKey = redact(redacted.OpenAI.APIKey)
		redacted.Local.APIKey = redact(redacted.Local.APIKey)
		data, _ := json.MarshalIndent(redacted, "", "  ")
		fmt.Println(string(data))
		if p, err := config.Path(); err == nil {
			fmt.Fprintf(os.Stderr, "\nConfig file: %s\n", p)
		}
	case "path":
		p, err := config.Path()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(p)
	default:
		fmt.Fprintf(os.Stderr, "unknown config command %q\n\n", args[0])
		printConfigUsage(os.Stderr)
		os.Exit(1)
	}
}

func redact(s string) string {
	if s == "" {
		return ""
	}
	if len(s) <= 8 {
		return "***"
	}
	return s[:4] + "..." + s[len(s)-4:]
}

func printConfigUsage(w io.Writer) {
	fmt.Fprintln(w, `Usage:
  commit-ai config set <key> <value>     Update a configuration value
  commit-ai config show                  Print the current config (keys redacted)
  commit-ai config path                  Print the config file path
  commit-ai config help                  Show this help

Keys:
  provider              openai | claude-code | codex | local
  openai.api_key        OpenAI API key
  openai.model          OpenAI model (default: gpt-4o)
  local.base_url        OpenAI-compatible endpoint (e.g. http://localhost:11434/v1)
  local.model           Local model name (e.g. llama3)
  local.api_key         Optional API key for local endpoint
  claude_code.path      Path to claude binary (default: claude in PATH)
  claude_code.model     Optional model override passed as --model
  codex.path            Path to codex binary (default: codex in PATH)
  codex.model           Optional model override passed as --model

Examples:
  commit-ai config set provider claude-code
  commit-ai config set openai.api_key sk-...
  commit-ai config set local.base_url http://localhost:11434/v1`)
}
