package src

import (
	"fmt"
	"time"
)

type Runner struct {
	fileManager FileManagerInterface
	git         GitInterface
	utils       UtilsInterface
	viewBuilder ViewBuilderInterface
}

func NewRunner(fm FileManagerInterface, g GitInterface, u UtilsInterface, b ViewBuilderInterface) *Runner {
	return &Runner{
		fileManager: fm,
		git:         g,
		utils:       u,
		viewBuilder: b,
	}
}

func (r *Runner) Start() {
	rules := DefaultRules()
	styles := DefaultStyles()

	// Check if current dir has .git (is local repository)
	// This is the return early error.
	if !r.git.IsGitRepository() {
		r.utils.ExitWithError("Current directory is not a git repository")
	}

	// Init and setup
	// Create instance of FileManager and setup.
	// FileManager should create the following:
	//
	// ~/.kcommit
	// ~/.kcommit/.kcommit_history.json

	r.fileManager.BasicSetup()

	// Check for rules on current dir
	// It may find .kcommitrc or not (not mandatory)
	// In case current project does not have .kcommitrc it should use a default config (DefaultRules)
	// More about kcommitrc on README.md.

	hasCustomConfig, err := r.fileManager.CheckIfPathExists(KcommitRcFileName)
	if err != nil {
		r.utils.HandleError(err, "Failed load kcommitrc")
	}

	if hasCustomConfig {

		customConfigStr, err := r.fileManager.ReadFileContent(KcommitRcFileName)
		if err != nil {
			r.utils.HandleError(err, "Failed to read .kcommitrc. Check if the formmat ir correct")
		}

		customRules, err := ParseJSONContent[CommitRulesDTO](customConfigStr)
		if err != nil {
			r.utils.HandleError(err, "Failed to parse .kcommitrc")
		}

		rules = customRules
	}

	// It should fetch some basic info in order to continue.
	// - Get current dir name as project name
	// - Get current branch
	// Those information should be needed to set kcommit_history.

	currentProjName, err := r.fileManager.GetCurrentDirectoryName()
	if err != nil {
		r.utils.HandleError(err, "Failed to read current dir name")
	}

	currentBranchName, err := r.git.GetCurrentBranch()
	if err != nil {
		r.utils.HandleError(err, "Failed to get current branch")
	}

	// Get the history content. If it's empty just set a basic structure.
	// kcommit_history should not be empty at this point.
	// It loads history to find the current scope to use on the commit.
	// If scope is empty it should prompt for user to set one.

	historyStr, err := r.fileManager.GetHistoryContent()
	if err != nil {
		r.utils.HandleError(err, "Failed to read kcommit history")
	}

	historyObj := &HistoryDTO{
		Projects: []ProjectDTO{
			{
				Name: currentProjName,
				Branches: []BranchDTO{
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
		h, err := ParseJSONContent[HistoryDTO](historyStr)
		if err != nil {
			r.utils.HandleError(err, "Failed to parse kcommit_history")
		}
		historyObj = h
	}

	history := historyObj.ToModel()

	// Check if history has project/branch.
	// Add project/branch to current history if needed.
	if !history.HasBranch(currentProjName, currentBranchName) {
		history.AddBranch(currentProjName, currentBranchName)
	}

	branchData, err := history.FindBranchData(currentProjName, currentBranchName)
	if err != nil {
		r.utils.HandleError(err, "Failed to locate project data")
	}

	// Define scope for current branch in case it's empty
	if branchData.Scope == "" {
		choices := []ListItem{
			{
				T: "branch",
				D: "use branch name as scope",
			},
			{
				T: "custom",
				D: "write a custom string to be the scope",
			},
		}

		answer := r.viewBuilder.NewListView("This branch does not have scope defined yet.", choices, 16)
		r.utils.ValidateInput(answer.T)

		if answer.T == "branch" {
			branchData.Scope = currentBranchName
		} else {
			newValue := r.viewBuilder.NewTextFieldView("Write a name for the scope", "")
			r.utils.ValidateInput(newValue)
			branchData.Scope = newValue
		}

	}

	// This will set the scope to be saved and the time it was updated.
	// Time updated is also used later to clear out old branches
	history.SetBranch(currentProjName, currentBranchName, branchData.Scope)

	// Choose commit type

	commitTypeOptions := r.utils.CommitTypeDTOsToListItems(rules.CommitTypeDTOs)
	selectCommitType := r.viewBuilder.NewListView("Please choose a commit type", commitTypeOptions, 32)
	r.utils.ValidateInput(selectCommitType.T)

	// Write commit message

	commitDescription := r.viewBuilder.NewTextFieldView("Write the commit message", "")
	r.utils.ValidateInput(commitDescription)

	// Build commit message

	commitMsg := fmt.Sprintf(
		"%s(%s): %s",
		selectCommitType.T,
		branchData.Scope,
		commitDescription,
	)

	// offer to commit of just print the commit message

	choices := []ListItem{
		{
			T: "commit",
			D: fmt.Sprintf("kcommit will call git commit with: %s", commitMsg),
		},
		{
			T: "just print",
			D: "kcommit will not call git commit, just print the resulting commit message",
		},
	}

	answer := r.viewBuilder.NewListView("This branch does not have scope defined yet.", choices, 16)
	r.utils.ValidateInput(answer.T)

	if answer.T == "commit" {
		msg, err := r.git.GitCommit(commitMsg)
		if err != nil {
			r.utils.HandleError(err, "Failed git commit")
		}
		println(styles.Text(msg, styles.AquamarineColor))
	} else {
		println(styles.Text(commitMsg, styles.AquamarineColor))
	}

	// Clean cache.
	history.CleanOldBranches(time.Now())

	// Save cleaned history.

	h, err := history.ToJson()
	if err != nil {
		r.utils.HandleError(err, "Failed to write history.json")
	}

	r.fileManager.WriteHistoryContent(h)
}
