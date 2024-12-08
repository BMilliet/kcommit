package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
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
			Title: commitType.Type,
			Desc:  commitType.Description,
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

func CreateView(model tea.Model) {
	a := tea.NewProgram(model)
	if _, err := a.Run(); err != nil {
		fmt.Printf("there's been an error: %v", err)
		os.Exit(1)
	}
}
