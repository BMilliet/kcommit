package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ParseJSONContent[T any](jsonString string) (*T, error) {
	var targetStruct T
	err := json.Unmarshal([]byte(jsonString), &targetStruct)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	return &targetStruct, nil
}

func CommitTypesToListItems(commitTypes []CommitType) []ListItem {
	var listItems []ListItem
	for _, commitType := range commitTypes {
		listItems = append(listItems, ListItem{
			T: commitType.Type,
			D: commitType.Description,
		})
	}
	return listItems
}

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running git command: %v", err)
	}

	branch := out.String()
	return branch[:len(branch)-1], nil
}

func GitCommit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git commit failed: %v\nOutput: %s", err, string(output))
	}

	fmt.Printf("Git commit successful: %s\n", string(output))
	return nil
}

func GetCurrentDirectoryName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %v", err)
	}

	return filepath.Base(dir), nil
}

func HasGitDirectory() bool {
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

func ValidateInput(v string) {
	if v == ExitSignal {
		os.Exit(0)
	}
}
