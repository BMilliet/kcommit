package main

import (
	"fmt"
	"time"

	"kcommit/src"
)

func main() {

	rules := src.DefaultRules()
	styles := src.DefaultStyles()

	// Check if current dir has .git
	// This is the return early error.
	if !src.HasGitDirectory() {
		src.ExitWithError("Current directory is not a git repository")
	}

	// Init and setup
	// Create instance of FileManager and setup.
	// FileManager should create the following:
	//
	// ~/.kcommit
	// ~/.kcommit/.kcommit_history.json

	fileManager, err := src.NewFileManager()
	if err != nil {
		src.HandleError(err, "Failed to initialize FileManager")
	}

	fileManager.BasicSetup()

	// Check for rules on current dir
	// It may find .kcommitrc or not (not mandatory)
	// In case current project does not have .kcommitrc it should use a default config (DefaultRules)
	// More about kcommitrc on README.md.

	hasCustomConfig, err := fileManager.CheckIfPathExists(src.KcommitRcFileName)
	if err != nil {
		src.HandleError(err, "Failed load kcommitrc")
	}

	if hasCustomConfig {

		customConfigStr, err := fileManager.ReadFileContent(src.KcommitRcFileName)
		if err != nil {
			src.HandleError(err, "Failed to read .kcommitrc. Check if the formmat ir correct")
		}

		customRules, err := src.ParseJSONContent[src.CommitRules](customConfigStr)
		if err != nil {
			src.HandleError(err, "Failed to parse .kcommitrc")
		}

		rules = customRules
	}

	// It should fetch some basic info in order to continue.
	// - Get current dir name as project name
	// - Get current branch
	// Those information should be needed to set kcommit_history.

	currentProjName, err := src.GetCurrentDirectoryName()
	if err != nil {
		src.HandleError(err, "Failed to read current dir name")
	}

	currentBranchName, err := src.GetCurrentBranch()
	if err != nil {
		src.HandleError(err, "Failed to get current branch")
	}

	// Get the history content. If it's empty just set a basic structure.
	// kcommit_history should not be empty at this point.
	// It loads history to find the current scope to use on the commit.
	// If scope is empty it should prompt for user to set one.

	historyStr, err := fileManager.GetHistoryContent()
	if err != nil {
		src.HandleError(err, "Failed to read kcommit history")
	}

	historyObj := &src.HistoryDTO{
		Projects: []src.ProjectModel{
			{
				Name: currentProjName,
				Branches: []src.BranchModel{
					{
						Name:      currentBranchName,
						Scope:     "",
						UpdatedAt: time.Now(),
					},
				},
			},
		},
	}

	if !(historyStr == "") {
		h, err := src.ParseJSONContent[src.HistoryDTO](historyStr)
		if err != nil {
			src.HandleError(err, "Failed to parse kcommit_history")
		}
		historyObj = h
	}

	history := src.CreateHistoryModelFromDTO(historyObj)

	// Check history has project/branch.
	// Add project/branch to current history if needed.
	if !history.HasBranch(currentProjName, currentBranchName) {
		history.AddBranch(currentProjName, currentBranchName)
	}

	branchData, err := history.FindBranchData(currentProjName, currentBranchName)
	if err != nil {
		src.HandleError(err, "Failed to locate project data")
	}

	// Define scope for this branch in case it's empty
	if branchData.Scope == "" {
		choices := []src.ListItem{
			{
				T: "branch",
				D: "use branch name as scope",
			},
			{
				T: "custom",
				D: "write a custom string to be the scope",
			},
		}

		answer := ""
		src.ListView("This branch does not have scope defined yet.", choices, 16, &answer)
		src.ValidateInput(answer)

		if answer == "branch" {
			branchData.Scope = currentBranchName
		} else {
			newValue := ""
			src.TextFieldView("Write a name for the scope", "", &newValue)
			src.ValidateInput(newValue)
			branchData.Scope = newValue
		}

	}

	// This will set the scope to be save and the time it was updated.
	// Time updated is also used lated to clear out old branches
	history.SetBranch(currentProjName, currentBranchName, branchData.Scope)

	// Choose commit type

	commitTypeOptions := src.CommitTypesToListItems(rules.CommitTypes)
	selectCommitType := ""
	src.ListView("Please choose a commit type", commitTypeOptions, 32, &selectCommitType)
	src.ValidateInput(selectCommitType)

	// Write commit message

	commitDescription := ""
	src.TextFieldView("Write the commit message", "", &commitDescription)
	src.ValidateInput(commitDescription)

	// Build commit message

	commitMsg := fmt.Sprintf(
		"%s(%s): %s",
		selectCommitType,
		branchData.Scope,
		commitDescription,
	)

	// offer to commit of just print the commit message

	choices := []src.ListItem{
		{
			T: "commit",
			D: fmt.Sprintf("kcommit will call git commit with: %s", commitMsg),
		},
		{
			T: "just print",
			D: "kcommit will not call git commit, just print the resulting commit message",
		},
	}

	answer := ""
	src.ListView("This branch does not have scope defined yet.", choices, 16, &answer)
	src.ValidateInput(answer)

	if answer == "commit" {
		msg, err := src.GitCommit(commitMsg)
		if err != nil {
			src.HandleError(err, "Failed git commit")
		}
		println(styles.Text(msg, styles.AquamarineColor))
	} else {
		println(styles.Text(commitMsg, styles.AquamarineColor))
	}

	// Clean cache.
	history.CleanOldBranches()

	// Save history

	h, err := history.ToJson()
	if err != nil {
		src.HandleError(err, "Failed to write history.json")
	}

	fileManager.WriteHistoryContent(h)
}
