# Commit-AI

Commit-AI is a CLI tool designed to automate the generation of meaningful commit messages for your Git repositories. It analyzes your git diff and crafts commit messages that follow the Conventional Commits spec.

Supported providers:

- **openai** — OpenAI API (e.g. gpt-4o), needs an API key
- **claude-code** — your Claude Code subscription, via the `claude` CLI
- **codex** — your ChatGPT Codex subscription, via the `codex` CLI
- **local** — any OpenAI-compatible local endpoint (Ollama, LM Studio, llama.cpp, vLLM, …)

## Installation

### With Go

```bash
go install github.com/Jonath-z/commit-ai@latest
```

### Linux binary

```bash
curl https://raw.githubusercontent.com/Jonath-z/commit-ai/master/install.sh | sh
```

## Configuration

Settings are persisted to `~/.config/commit-ai/config.json`. You configure once, then just run `commit-ai`.

```bash
# Pick a provider
commit-ai config set provider openai          # or claude-code, codex, local

# Provider-specific settings
commit-ai config set openai.api_key sk-...
commit-ai config set openai.model gpt-4o

commit-ai config set local.base_url http://localhost:11434/v1
commit-ai config set local.model llama3

commit-ai config set claude_code.model sonnet     # optional
commit-ai config set codex.model gpt-5            # optional

# Inspect (API keys are redacted) or locate the config
commit-ai config show
commit-ai config path
```

Backward-compatible shortcut: `commit-ai --apiKey sk-...` saves the OpenAI key to config, sets `provider=openai`, and runs.

### Using the Claude Code subscription

1. Install and authenticate the [Claude Code CLI](https://docs.claude.com/en/docs/claude-code) (`claude login`).
2. `commit-ai config set provider claude-code`
3. Optional: `commit-ai config set claude_code.model sonnet`

The tool shells out to `claude -p`, so usage counts against your existing Claude subscription.

### Using the Codex subscription

1. Install and authenticate the [Codex CLI](https://github.com/openai/codex) with your ChatGPT account.
2. `commit-ai config set provider codex`
3. Optional: `commit-ai config set codex.model gpt-5`

The tool shells out to `codex exec`, so usage counts against your existing ChatGPT subscription.

### Using a local model

Run any OpenAI-compatible server. With Ollama:

```bash
ollama serve
ollama pull llama3
commit-ai config set provider local
commit-ai config set local.base_url http://localhost:11434/v1
commit-ai config set local.model llama3
```

## Usage

Inside a git repo with unstaged changes:

```bash
commit-ai
```

Override the configured provider for a single run:

```bash
commit-ai --provider local
```

## Contributing

If you would like to contribute to the project, please fork the repository and open a pull request.

## License

MIT — see the LICENSE file for details.
