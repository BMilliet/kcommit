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
	cr := src.DefaultRules()

	hasCustomConfig, err := fileManager.CheckIfPathExists(".kcommitrc")
	if err != nil {
		log.Fatal(err)
	}

	if hasCustomConfig {

		customConfigStr, err := fileManager.ReadFileContent(".kcommitrc")
		if err != nil {
			log.Fatalf("Failed to read .kcommitrc. Check if the formmat ir correct: %v", err)
		}

		customRules, err := src.ParseJSONContent[src.CommitRules](customConfigStr)
		if err != nil {
			log.Fatalf("Failed to read .kcommitrc. Check if the formmat ir correct: %v", err)
		}

		cr = customRules
	}

	// check if user have .kcommit/history

	// load header from history if present

	// define header if there is no history offer to use branch name

	// choose commit type

	commitTypeOptions := src.CommitTypesToListItems(cr.CommitTypes)

	p := tea.NewProgram(src.ListView("Please choose a commit type", commitTypeOptions))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
