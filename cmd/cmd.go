package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/Jonath-z/commit-ai/src"
	"github.com/Jonath-z/commit-ai/src/config"
	"github.com/Jonath-z/commit-ai/src/providers"
)

func Cli() {
	if len(os.Args) >= 2 && os.Args[1] == "config" {
		runConfig(os.Args[2:])
		return
	}

	if handleTopLevelHelp() {
		return
	}

	flag.CommandLine.SetOutput(os.Stderr)
	flag.Usage = func() { printRootHelp(os.Stderr) }
	apiKey := flag.String("apiKey", "", "OpenAI API key (saved to config; sets provider to openai)")
	provider := flag.String("provider", "", "Override provider for this run (openai, claude-code, codex, local)")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	if *apiKey != "" {
		cfg.OpenAI.APIKey = *apiKey
		cfg.Provider = "openai"
		if err := config.Save(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "failed to save config: %v\n", err)
			os.Exit(1)
		}
	}

	if *provider != "" {
		if !config.ValidProvider(*provider) {
			fmt.Fprintf(os.Stderr, "unknown provider %q (valid: openai, claude-code, codex, local)\n", *provider)
			os.Exit(1)
		}
		cfg.Provider = *provider
	}

	p, err := providers.New(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	gitDiff := src.GetGitChanges()
	spinner := src.StartSpinner("generating commit message")
	commitMsg, err := src.GenerateCommitMessage(p, gitDiff)
	spinner.Stop()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate commit message: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(commitMsg)
	fmt.Println("------------- Executing the commit message ------------")
	src.ExecuteCommitMsg(commitMsg)
}
