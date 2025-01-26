package src

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type GitInterface interface {
	GetCurrentBranch() (string, error)
	GitCommit(msg string) (string, error)
	IsGitRepository() bool
}

type Git struct{}

func NewGit() *Git {
	return &Git{}
}

func (g *Git) GetCurrentBranch() (string, error) {
	branch, err := g.execGitCommand("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return "", fmt.Errorf("GetCurrentBranch -> %v", err)
	}
	return branch, nil
}

func (g *Git) GitCommit(msg string) (string, error) {
	output, err := g.execGitCommand("commit", "-m", msg)
	if err != nil {
		return "", fmt.Errorf("GitCommit -> %v", err)
	}
	return output, nil
}

func (g *Git) IsGitRepository() bool {
	currentDir, err := os.Getwd()
	if err != nil {
		return false
	}

	gitDir := filepath.Join(currentDir, ".git")

	info, err := os.Stat(gitDir)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func (g *Git) execGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s", strings.TrimSpace(stdout.String()))
	}

	return strings.TrimSpace(stdout.String()), nil
}
