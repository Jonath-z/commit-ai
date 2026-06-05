package cmd

import (
	"io"
	"os"
)

const rootHelp = `commit-ai — generate git commit messages with AI

Usage:
  commit-ai [flags]              Generate a commit message and run git commit
  commit-ai config <command>     Manage persisted configuration
  commit-ai help                 Show this help

Providers (select with ` + "`commit-ai config set provider <name>`" + `):
  openai         OpenAI API (needs API key)
  claude-code    Claude Code subscription via the ` + "`claude`" + ` CLI
  codex          ChatGPT Codex subscription via the ` + "`codex`" + ` CLI
  local          OpenAI-compatible local endpoint (Ollama, LM Studio, ...)

Flags:
  --apiKey <key>       Save the OpenAI API key to config and run (provider=openai)
  --provider <name>    Override the configured provider for this run
  -h, --help           Show this help

Examples:
  commit-ai                                       # use saved config
  commit-ai --provider local                      # override provider once
  commit-ai --apiKey sk-...                       # first-time OpenAI setup
  commit-ai config set provider claude-code       # switch to Claude Code
  commit-ai config set openai.api_key sk-...      # save the API key once
  commit-ai config show                           # inspect current config

Config file: ~/.config/commit-ai/config.json (see ` + "`commit-ai config path`" + `)

Run ` + "`commit-ai config`" + ` for configuration command details.
`

func isHelpArg(s string) bool {
	switch s {
	case "-h", "--help", "help":
		return true
	}
	return false
}

func printRootHelp(w io.Writer) {
	_, _ = io.WriteString(w, rootHelp)
}

func handleTopLevelHelp() bool {
	for _, a := range os.Args[1:] {
		if a == "--" {
			return false
		}
		if isHelpArg(a) {
			printRootHelp(os.Stdout)
			return true
		}
	}
	return false
}
