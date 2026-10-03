# Commit-AI v1.1.0 — Bring Your Own Model, Commit One File at a Time

Commit-AI started with a simple idea: your git diff already explains what you did — you shouldn't have to write it up a second time. You run `commit-ai`, it reads your changes, and it writes a [Conventional Commits](https://www.conventionalcommits.org/) message for you.

Until now, that came with two big assumptions: you had an OpenAI API key, and you wanted to commit *everything* in your working tree. Version 1.1.0 removes both, and makes keeping the tool up to date a one-liner.

## Use the AI subscription you already pay for

Commit-AI is no longer tied to the OpenAI API. It now supports four providers, and switching is a single config command:

```bash
commit-ai config set provider claude-code    # Claude subscription via the claude CLI
commit-ai config set provider codex          # ChatGPT subscription via the codex CLI
commit-ai config set provider local          # Ollama, LM Studio, any OpenAI-compatible endpoint
commit-ai config set provider openai         # the classic API-key route
```

If you already use Claude Code or Codex, there is nothing new to pay for and no API key to manage — commit-ai shells out to the CLI you already authenticated, so usage counts against your existing subscription.

And if you want your diffs to never leave your machine, point it at a local model:

```bash
commit-ai config set provider local
commit-ai config set local.base_url http://localhost:11434/v1
commit-ai config set local.model llama3
```

Your choice is saved to `~/.config/commit-ai/config.json`, and you can override it for a single run with `--provider`:

```bash
commit-ai --provider local
```

## Commit just one file

Previously, commit-ai staged your whole working tree (`git add .`) and committed everything in one go. That's fine at the end of a focused session — and wrong the moment your tree mixes two unrelated changes.

Now you can pass paths, and commit-ai scopes the entire flow to them:

```bash
commit-ai src/parser.go              # one file
commit-ai src/parser.go README.md    # a few files
commit-ai src/                       # a directory (git pathspec rules apply)
```

Only those files are diffed, so the generated message describes exactly what you're committing — and only those files are staged, so the rest of your changes stay untouched for the next commit. Run it with no arguments and it behaves exactly as before.

## Self-updating

Installing was already easy; staying current wasn't. Two new commands fix that:

```bash
commit-ai version    # what you have
commit-ai update     # get the latest release
```

`update` checks the latest GitHub release, downloads the right binary for your OS and architecture, and swaps it in atomically — a failed download can never leave you with a broken install. If your binary lives in a root-owned directory like `/usr/local/bin`, run `sudo commit-ai update`.

## The small stuff

- A spinner while your message is being generated, so the CLI no longer sits silently on a slow model.
- A real help system: `commit-ai help` and `commit-ai config` document every command and flag.
- Fixes to the install script and to API-key persistence.

## Getting started

New install:

```bash
curl https://raw.githubusercontent.com/Jonath-z/commit-ai/master/install.sh | sh
```

Or with Go:

```bash
go install github.com/Jonath-z/commit-ai@latest
```

Already on an older version? Re-run the install script once — every release after that is just `commit-ai update`.

Then, inside any repo with changes:

```bash
commit-ai
```

That's it. Your diff goes in, a Conventional Commits message comes out, and the commit is made.

---

Commit-AI is open source under the MIT license. Issues and pull requests are welcome at [github.com/Jonath-z/commit-ai](https://github.com/Jonath-z/commit-ai).
