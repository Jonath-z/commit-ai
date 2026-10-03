package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const repo = "Jonath-z/commit-ai"

func runUpdate() {
	latest, err := latestReleaseTag()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to check latest release: %v\n", err)
		os.Exit(1)
	}

	current := "v" + Version
	if latest == current {
		fmt.Printf("commit-ai is already up to date (%s)\n", current)
		return
	}

	fmt.Printf("updating commit-ai %s -> %s\n", current, latest)

	asset := fmt.Sprintf("commit-ai-%s-%s", runtime.GOOS, runtime.GOARCH)
	url := fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", repo, latest, asset)

	exe, err := os.Executable()
	if err == nil {
		exe, err = filepath.EvalSymlinks(exe)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to locate current binary: %v\n", err)
		os.Exit(1)
	}

	if err := replaceBinary(url, exe); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if os.IsPermission(err) || strings.Contains(err.Error(), "permission denied") {
			fmt.Fprintln(os.Stderr, "no write access to", filepath.Dir(exe), "- try: sudo commit-ai update")
		}
		os.Exit(1)
	}

	fmt.Printf("commit-ai updated to %s\n", latest)
}

func latestReleaseTag() (string, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get("https://api.github.com/repos/" + repo + "/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned %s", resp.Status)
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}
	if release.TagName == "" {
		return "", fmt.Errorf("no release tag found")
	}
	return release.TagName, nil
}

func replaceBinary(url, exe string) error {
	client := &http.Client{Timeout: 5 * time.Minute}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s returned %s", url, resp.Status)
	}

	// Download next to the current binary so the final rename is atomic
	// and stays on the same filesystem.
	tmp, err := os.CreateTemp(filepath.Dir(exe), ".commit-ai-update-*")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)

	if _, err := io.Copy(tmp, resp.Body); err != nil {
		tmp.Close()
		return fmt.Errorf("download failed: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(tmpPath, 0o755); err != nil {
		return err
	}
	return os.Rename(tmpPath, exe)
}
