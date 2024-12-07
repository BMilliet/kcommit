package main

import (
	"fmt"
	"log"
	"os"

	"kcommit/src"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	// check if user have .kcommit/ on home

	fileManager, err := src.NewFileManager()
	if err != nil {
		log.Fatalf("Failed to initialize FileManager: %v", err)
	}

	fileManager.BasicSetup()

	// check if current project has .kcommitrc for custom list else use default karma list
	l := []string{
		"feat",
		"fix",
		"chore",
		"refactor",
		"style",
		"test",
		"docs",
	}

	hasCustomConfig, err := fileManager.CheckIfPathExists(".kcommitrc")
	if err != nil {
		log.Fatal(err)
	}

	if hasCustomConfig {

		customConfigContent, err := fileManager.ReadFileContent(".kcommitrc")
		if err != nil {
			log.Fatalf("Failed to read .kcommitrc. Check if the formmat ir correct: %v", err)
		}

		print(customConfigContent)

		// read and parse kcommitrc
	}

	// check if user have .kcommit/history

	// load header from history if present

	// define header if there is no history offer to use branch name

	// choose commit type

	p := tea.NewProgram(src.ListView("Please choose a commit type", l))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
