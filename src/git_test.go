package src

import (
	"os"
	"os/exec"
	"testing"
)

func TestGitGetCurrentBranchFromUnbornRepo(t *testing.T) {
	tempDir := t.TempDir()
	runGit(t, tempDir, "init")
	runGit(t, tempDir, "checkout", "-b", "first-branch")

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current directory: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(originalDir)
	})

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}

	branch, err := NewGit().GetCurrentBranch()
	if err != nil {
		t.Fatalf("expected branch name, got error: %v", err)
	}

	if branch != "first-branch" {
		t.Fatalf("expected branch first-branch, got %q", branch)
	}
}

func runGit(t *testing.T, dir string, args ...string) {
	t.Helper()

	cmd := exec.Command("git", args...)
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v failed: %v\n%s", args, err, output)
	}
}
